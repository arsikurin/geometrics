package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

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
				return errors.WithMessage(err, "check whether problem exists failed in patch problem by id")
			}

			return echo.ErrNotFound
		}

		ppr := new(types.POSTProblemReq)
		if err := c.Bind(ppr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(ppr); err != nil {
			return errors.WithMessage(err, "validation failed in post problem by id")
		}

		res, err := utils.CheckProblem(strconv.Itoa(id), ppr.GgbBase64)
		if err != nil {
			c.Logger().Error(err)
		}

		go utils.InsertSubmit(ctx, c, id, int(res), ppr.GgbBase64)

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":   http.StatusOK,
			"status": "ok",
			"result": res,
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
			return errors.WithMessage(err, "validation failed in put problem")
		}

		problem := models.Problem{
			Name:        ppr.Name,
			Description: ppr.Description,
			SolutionRaw: ppr.SolutionBase64,
		}

		err := problem.InsertG(ctx, boil.Infer())
		if err != nil {
			return errors.WithMessage(err, "insert problem failed in put problem by id")
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
			return errors.WithMessage(err, "validation failed in patch problem by id")
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

		problemID, err := problem.DeleteG(ctx)
		if err != nil {
			return errors.WithMessage(err, "delete problem failed in delete problem by id")
		}

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":       http.StatusOK,
			"status":     "ok",
			"problem_id": problemID,
		})
	}
}

func Login(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		lcr := new(types.LoginReq)
		if err := c.Bind(lcr); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(lcr); err != nil {
			return errors.WithMessage(err, "validation failed in login")
		}

		// username := c.FormValue("login")
		// password := c.FormValue("password")
		if isExists, err := models.Users(Where("login=?", lcr.Login)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in login")
			}

			return c.JSON(http.StatusUnauthorized, echo.Map{ //nolint:wrapcheck
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

		return c.JSON(http.StatusOK, echo.Map{ //nolint:wrapcheck
			"code":   http.StatusOK,
			"status": "ok",
			"token":  t,
		})
	}
}
