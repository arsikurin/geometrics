package utils

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"

	"geometrics/auth"
	"geometrics/types"
)

func AuthMiddleware(AllowNoToken bool) echo.MiddlewareFunc {
	ErrorHandlerWithContext := func(err error, c echo.Context) error { return err }
	if AllowNoToken {
		ErrorHandlerWithContext = func(err error, c echo.Context) error {
			if _, ok := err.(*echo.HTTPError); ok {
				c.Set("user", &jwt.Token{
					Method: jwt.SigningMethodRS256,
					Claims: &auth.JWTCustomClaims{
						UserID:  -1,
						Name:    "Unauthorized",
						IsAdmin: false,
						StandardClaims: jwt.StandardClaims{
							ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //nolint:gomnd
						},
					},
				})

				return nil
			}

			return err
		}
	}

	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &auth.JWTCustomClaims{},
		SigningKey: func() *rsa.PublicKey {
			key, err := auth.GetRSAPublicKey()
			if err != nil {
				log.Println(err)
			}

			return key
		}(),
		ContinueOnIgnoredError:  AllowNoToken,
		ErrorHandlerWithContext: ErrorHandlerWithContext,
		TokenLookup:             "header:Authorization,cookie:token",
		SigningMethod:           "RS256",
	})
}

func LoggerMiddleware() echo.MiddlewareFunc {
	logger := zerolog.New(os.Stdout)

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("method", v.Method).
				Time("time", v.StartTime).
				Err(v.Error).
				Msg("request")

			return nil
		},
	})
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
