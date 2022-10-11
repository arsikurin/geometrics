package api

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"geometrics/auth"
	"geometrics/models"
	"geometrics/types"
)

func GETProblemByID(c echo.Context) error {
	id := c.Param("id")

	action := c.QueryParam("action")
	if action == "" {
		action = "view"
	}

	return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/problems/%s?action=%s", id, action))
}
func POSTProblemByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		ppr := new(types.POSTProblemReq)
		if err := c.Bind(ppr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(ppr); err != nil {
			return err
		}

		var out, scriptErr bytes.Buffer

		cmd := exec.Command("/usr/local/bin/python", "main.py", strconv.Itoa(id), ppr.GgbBase64)
		cmd.Stdout = &out
		cmd.Stderr = &scriptErr

		err = cmd.Run()
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("python checker failed: %s", scriptErr.String()))
		}

		res, err := strconv.Atoi(out.String())
		if err != nil {
			return errors.WithMessage(err, "not int response from python checker")
		}

		go func() {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*auth.JWTCustomClaims)

			var newSubmit models.Submit
			newSubmit.UserID = claims.UserID
			newSubmit.ProblemID = id
			newSubmit.Status = res
			newSubmit.SolutionRaw = ppr.GgbBase64

			err = newSubmit.InsertG(ctx, boil.Infer())
			if err != nil {
				c.Logger().Error(errors.WithMessage(err, "insert submit failed in post problem by id"))
			}
		}()

		return c.JSON(http.StatusOK, echo.Map{
			"code":   http.StatusOK,
			"status": "ok",
			"result": res,
		})
	}
}
func PUTProblemByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		ppr := new(types.PUTProblemReq)
		if err := c.Bind(ppr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(ppr); err != nil {
			return err
		}

		problem, err := models.Problems(SQL("INSERT INTO problems (name, description, solution_raw) VALUES ($1, $2, $3) RETURNING id", ppr.Name, ppr.Description, ppr.SolutionBase64)).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "insert problem failed in put problem by id")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"code":       http.StatusOK,
			"status":     "ok",
			"problem_id": problem.ID,
		})
	}
}
func PATCHProblemByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		if isExists, err := models.Problems(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether problem exists failed in patch problem by id")
			}

			return echo.ErrNotFound
		}

		ppr := new(types.PATCHProblemReq)
		if err := c.Bind(ppr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(ppr); err != nil {
			return err
		}

		problem, err := models.FindProblemG(ctx, id)
		problem.Name = ppr.Name
		problem.Description = ppr.Description
		problem.SolutionRaw = ppr.SolutionBase64

		_, err = problem.UpdateG(ctx, boil.Infer())
		if err != nil {
			return errors.WithMessage(err, "update problem failed in patch problem by id")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"code":   http.StatusOK,
			"status": "ok",
		})
	}
}
func DELETEProblemByID(c echo.Context) error {
	id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTCustomClaims)
	name := claims.Name

	return c.String(http.StatusOK, "method DELETE "+id+" "+name)
}

func Login(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		lcr := new(types.LoginReq)
		if err := c.Bind(lcr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(lcr); err != nil {
			return err
		}

		// username := c.FormValue("login")
		// password := c.FormValue("password")
		if isExists, err := models.Users(Where("login=?", lcr.Login)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in login")
			}

			return c.JSON(http.StatusUnauthorized, echo.Map{
				"code":   http.StatusUnauthorized,
				"status": "error",
				"message": fmt.Sprintf(
					"%d %s", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized),
				),
				"detail": "invalid credentials",
			})
		}

		user, err := models.Users(Where("login=?", lcr.Login)).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get user from the db failed in login")
		}

		if lcr.Password != user.Password {
			return echo.ErrUnauthorized
		}

		t, err := auth.GenerateAccessToken(user)
		if err != nil {
			return errors.WithMessage(err, "generate access token failed in login")
		}

		c.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    t,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 24), //nolint:gomnd
			Secure:   false,
			HttpOnly: true,
		})

		return c.JSON(http.StatusOK, echo.Map{
			"code":   http.StatusOK,
			"status": "ok",
			"token":  t,
		})
	}
}
