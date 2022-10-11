package profiles

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"geometrics/auth"
	"geometrics/models"
)

func GETProfileByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		if isExists, err := models.Users(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in get profile by id")
			}

			return echo.ErrNotFound
		}

		user, err := models.Users(Where("id=?", id)).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get user from the db failed in get profile by id")
		}

		return c.Render(http.StatusOK, "profile.gohtml", map[string]interface{}{
			"name": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			"user": user,
		})
	}
}

func GETProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTCustomClaims)

	if claims.UserID != -1 {
		return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/profiles/%d", claims.UserID))
	}

	return c.Redirect(http.StatusPermanentRedirect, "/login")
}
