package submits

import (
	"context"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"geometrics/auth"
	"geometrics/models"
)

func GETSubmitByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		if isExists, err := models.Submits(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether submit exists failed in get submit by id")
			}

			return echo.ErrNotFound
		}

		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("assert token failed in get submit by id")
		}

		claims, ok := user.Claims.(*auth.JWTCustomClaims)
		if !ok {
			return errors.New("assert claims failed in get submit by id")
		}

		submit, err := models.Submits(
			Select(models.SubmitColumns.UserID, models.SubmitColumns.SolutionRaw),
			models.SubmitWhere.ID.EQ(id),
		).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get submit failed in get submit by id")
		}

		if claims.UserID != submit.UserID {
			return echo.ErrForbidden
		}

		return c.Render(http.StatusOK, "submit.gohtml", map[string]interface{}{ //nolint:wrapcheck
			"submit": submit,
		})
	}
}
