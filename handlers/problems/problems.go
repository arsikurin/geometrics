package problems

import (
	"context"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"geometrics/auth"
	"geometrics/models"
)

func GETProblemByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		if isExists, err := models.Problems(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in get problem by id")
			}

			return echo.ErrNotFound
		}

		problem, err := models.Problems(
			Select(models.ProblemTableColumns.ID, models.ProblemColumns.Name, models.ProblemColumns.Description),
			models.ProblemWhere.ID.EQ(id),
		).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get problem failed in get problem by id")
		}

		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("assert token failed in get problem by id")
		}

		claims, ok := user.Claims.(*auth.JWTCustomClaims)
		if !ok {
			return errors.New("assert claims failed in get problem by id")
		}

		var submits models.SubmitSlice
		if claims.UserID != -1 {
			submits, err = models.Submits(
				Select(models.SubmitColumns.ID, models.SubmitColumns.Status, models.SubmitColumns.CreatedAt),
				models.SubmitWhere.UserID.EQ(claims.UserID),
				models.SubmitWhere.ProblemID.EQ(problem.ID),
				OrderBy(models.SubmitColumns.ID),
				Limit(20), //nolint:gomnd
			).AllG(ctx)
			if err != nil {
				return errors.WithMessage(err, "get submits failed in get problem by id")
			}
		}

		return c.Render(http.StatusOK, "problem.gohtml", map[string]interface{}{ //nolint:wrapcheck
			"submits": submits,
			"problem": problem,
		})
	}
}

func GETSubmitsByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		if isExists, err := models.Problems(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether problem exists failed in get submits by id")
			}

			return echo.ErrNotFound
		}

		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("assert token failed in get submits by id")
		}

		claims, ok := user.Claims.(*auth.JWTCustomClaims)
		if !ok {
			return errors.New("assert claims failed in get submits by id")
		}

		problem, err := models.Problems(
			Select(models.ProblemTableColumns.ID, models.ProblemColumns.Name, models.ProblemColumns.Description),
			models.ProblemWhere.ID.EQ(id),
		).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get problem failed in get submits by id")
		}

		submits, err := models.Submits(
			Select(models.SubmitColumns.ID, models.SubmitColumns.Status, models.SubmitColumns.CreatedAt, models.SubmitColumns.UserID),
			models.SubmitWhere.ProblemID.EQ(problem.ID),
			OrderBy(models.SubmitColumns.ID),
			Limit(20), //nolint:gomnd
		).AllG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get submits failed in get submits by id")
		}

		return c.Render(http.StatusOK, "submits.gohtml", map[string]interface{}{ //nolint:wrapcheck
			"submits": submits,
			"problem": problem,
			"current": claims.UserID,
		})
	}
}

func GETSolveByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		if isExists, err := models.Problems(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in get solve by id")
			}

			return echo.ErrNotFound
		}

		problem, err := models.Problems(
			models.ProblemWhere.ID.EQ(id),
		).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get problem failed in get solve by id")
		}

		return c.Render(http.StatusOK, "solve.gohtml", map[string]interface{}{ //nolint:wrapcheck
			"problem": problem,
		})
	}
}
