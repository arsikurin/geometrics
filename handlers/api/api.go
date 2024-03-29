package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"

	"geometrics/auth"
	"geometrics/models"
	"geometrics/types"
	"geometrics/utils"
)

func GETProblemByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/problems/%d", id)) //nolint:wrapcheck
}

func POSTProblemByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		if isExists, err := models.Problems(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether problem exists failed in post problem by id")
			}

			return echo.ErrNotFound
		}

		ppr := new(types.POSTProblemReq)
		if err := c.Bind(ppr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(ppr); err != nil {
			return err
		}

		res, err := utils.CheckProblem(strconv.Itoa(id), ppr.GgbBase64)
		if err != nil {
			c.Logger().Error(err)
		}

		go utils.InsertSubmit(ctx, c, id, res, ppr.GgbBase64)

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":   http.StatusOK,
			"status": "ok",
			"result": res,
			"detail": res.String(),
		})
	}
}

func PUTProblem(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		ppr := new(types.PUTProblemReq)
		if err := c.Bind(ppr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(ppr); err != nil {
			return err
		}

		problem := models.Problem{
			Name:        ppr.Name,
			Description: ppr.Description,
			SolutionRaw: ppr.SolutionBase64,
		}

		err := problem.InsertG(ctx, boil.Infer())
		if err != nil {
			return errors.WithMessage(err, "insert problem failed in put problem")
		}

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
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
		if err != nil {
			return errors.WithMessage(err, "find problem failed in patch problem by id")
		}

		problem.Name = ppr.Name
		problem.Description = ppr.Description
		problem.SolutionRaw = ppr.SolutionBase64

		problemID, err := problem.UpdateG(ctx, boil.Infer())
		if err != nil {
			return errors.WithMessage(err, "update problem failed in patch problem by id")
		}

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":       http.StatusOK,
			"status":     "ok",
			"problem_id": problemID,
		})
	}
}

func DELETEProblemByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("assert token failed in delete problem by id")
		}

		claims, ok := user.Claims.(*auth.JWTCustomClaims)
		if !ok {
			return errors.New("assert claims failed in delete problem by id")
		}

		if !claims.IsAdmin {
			return echo.ErrForbidden
		}

		if isExists, err := models.Problems(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether problem exists failed in delete problem by id")
			}

			return echo.ErrNotFound
		}

		problem, err := models.FindProblemG(ctx, id)
		if err != nil {
			return errors.WithMessage(err, "find problem failed in delete problem by id")
		}

		_, err = problem.DeleteG(ctx)
		if err != nil {
			return errors.WithMessage(err, "delete problem failed in delete problem by id")
		}

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":   http.StatusOK,
			"status": "ok",
		})
	}
}

func POSTLogin(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		lr := new(types.LoginReq)
		if err := c.Bind(lr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(lr); err != nil {
			return err
		}

		if isExists, err := models.Users(Where("login=?", lr.Login)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in login")
			}

			return echo.ErrUnauthorized
		}

		user, err := models.Users(Where("login=?", lr.Login)).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get user from the db failed in login")
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(lr.Password)); err != nil {
			return echo.ErrUnauthorized
		}

		go func() {
			user.LastOnline = time.Now()

			_, err := user.UpdateG(ctx, boil.Infer())
			if err != nil {
				c.Logger().Error(errors.WithMessage(err, "update user last online failed in login"))
			}
		}()
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

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":   http.StatusOK,
			"status": "ok",
			"token":  t,
		})
	}
}

func POSTRegister(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		rr := new(types.RegisterReq)
		if err := c.Bind(rr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(rr); err != nil {
			return err
		}

		if isExists, err := models.Users(Where("login=?", rr.Login)).ExistsG(ctx); isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in login")
			}

			return c.JSON(http.StatusUnauthorized, echo.Map{ //nolint:wrapcheck
				"code":   http.StatusUnauthorized,
				"status": "error",
				"message": fmt.Sprintf(
					"%d %s", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized),
				),
				"detail": "пользователь существует",
			})
		}
		userGrade, err := strconv.Atoi(rr.Grade)
		if err != nil {
			userGrade = 0
			c.Logger().Info(errors.WithMessage(err, "conversion grade to int failed in register"))
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rr.Password), 13)
		if err != nil {
			return errors.WithMessage(err, "hash password failed in register")
		}

		user := models.User{
			Login:     rr.Login,
			Password:  hashedPassword,
			Type:      int(types.Student),
			FirstName: rr.FirstName,
			LastName:  rr.LastName,
			Grade: null.Int{
				Int:   userGrade,
				Valid: func() bool { return userGrade != 0 }(),
			},
			School: null.String{
				String: rr.School,
				Valid:  func() bool { return rr.School != "" }(),
			},
		}

		err = user.InsertG(ctx, boil.Infer())
		if err != nil {
			return errors.WithMessage(err, "insert user failed in register")
		}

		t, err := auth.GenerateAccessToken(&user)
		if err != nil {
			return errors.WithMessage(err, "generate access token failed in register")
		}

		c.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    t,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 24), //nolint:gomnd
			Secure:   false,
			HttpOnly: true,
		})

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":   http.StatusOK,
			"status": "ok",
			"token":  t,
		})
	}
}
