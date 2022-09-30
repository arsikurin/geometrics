package profiles

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

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
		claims := c.Get("user").(*jwt.Token).Claims.(*auth.JWTCustomClaims)
		name := claims.Name

		if isExists, err := models.Users(Where("id=?", claims.UserID)).ExistsG(ctx); !isExists {
			if err != nil {
				return err
			}
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"code":   http.StatusUnauthorized,
				"status": "error",
				"message": fmt.Sprintf(
					"%d %s", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized),
				),
				"detail": "user not exists",
			})
		}
		user, err := models.Users(Where("id=?", claims.UserID)).OneG(ctx)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, fmt.Sprintf("profile %d %s %d %v %v %s %s", id, name, user.Type, user.Grade, user.School, user.CreatedAt, user.Timezone))
	}
}

func GETProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTCustomClaims)

	return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/profiles/%d", claims.UserID))
}
