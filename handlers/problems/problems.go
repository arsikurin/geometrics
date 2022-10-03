package problems

import (
	"context"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"geometrics/auth"
)

func GETProblemByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		action := c.QueryParam("action")
		if action == "" {
			action = "view"
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.JWTCustomClaims)
		name := claims.Name

		return c.String(http.StatusOK, "problem "+strconv.Itoa(id)+" "+name)
	}
}
