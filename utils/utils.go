package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"geometrics/auth"
	"geometrics/types"
)

func UseAuth(group *echo.Group) {
	key, err := auth.GetRSAPublicKey()
	if err != nil {
		log.Println(err)
	}

	group.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &auth.JWTCustomClaims{},
		SigningKey:    key,
		TokenLookup:   "header:Authorization,cookie:token",
		SigningMethod: "RS256",
	}))
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	title := fmt.Sprintf("%d Internal Server Error", code)
	var detail []types.APIError

	if he, ok := err.(*echo.HTTPError); ok {
		if detail, ok = he.Message.([]types.APIError); ok {
			code = he.Code
			title = fmt.Sprintf("%d Validation Failed", code)
		} else if message, ok := he.Message.(string); ok && message == "missing or malformed jwt" {
			code = http.StatusForbidden
			title = fmt.Sprintf("%d %s", code, http.StatusText(code))
		} else {
			code = he.Code
			title = fmt.Sprintf("%d %s", code, http.StatusText(code))
		}
	}

	if strings.Split(c.Path(), "/")[1] == "api" {
		if err := c.JSON(code, map[string]interface{}{
			"code":    code,
			"status":  "error",
			"message": title,
			"detail":  detail,
		}); err != nil {
			c.Logger().Error(err)
		}
	} else {
		if err := c.Render(code, "error.html", map[string]interface{}{
			"error": title,
			// "detail": detail,
		}); err != nil {
			c.Logger().Error(err)
		}
	}
	c.Logger().Error(err)
}
