package utils

import (
	"bytes"
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"geometrics/auth"
	"geometrics/models"
	"geometrics/types"
)

func AuthMiddleware(allowNoToken bool) echo.MiddlewareFunc {
	ErrorHandlerWithContext := func(err error, c echo.Context) error { return err }
	if allowNoToken {
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
				c.Logger().Error(errors.WithMessage(err, "auth middleware"))

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
		ContinueOnIgnoredError:  allowNoToken,
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

	var (
		detail             interface{}
		JWTValidationError *jwt.ValidationError
		link               = "/"
		linkT              = "click here to go home"
		fix                string
	)

	if HTTPErr, ok := err.(*echo.HTTPError); ok {
		if detail, ok = HTTPErr.Message.([]types.APIError); ok {
			code = HTTPErr.Code
			title = fmt.Sprintf("%d Validation Failed", code)
		} else if message, ok := HTTPErr.Message.(string); ok && message == "missing or malformed jwt" {
			code = http.StatusForbidden
			detail = "Consider logging in"
			title = fmt.Sprintf("%d %s", code, http.StatusText(code))
		} else {
			code = HTTPErr.Code
			title = fmt.Sprintf("%d %s", code, http.StatusText(code))
		}
	} else if errors.As(err, &JWTValidationError) {
		code = http.StatusForbidden
		title = fmt.Sprintf("%d %s", code, http.StatusText(code))
		detail = "Consider logging in"
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
		if code == http.StatusForbidden || code == http.StatusUnauthorized {
			err := c.Redirect(http.StatusFound, "/login") //nolint:wrapcheck
			if err != nil {
				c.Logger().Error(err)
			}

		} else if err := c.Render(code, "error.gohtml", map[string]interface{}{
			"error":     title,
			"detail":    detail,
			"link":      link,
			"link_text": linkT,
			"fix":       fix,
		}); err != nil {
			c.Logger().Error(err)
		}
	}

	c.Logger().Error(err)
}

func CheckProblem(id string, solution string) (types.ProblemResult, error) {
	var out, scriptErr bytes.Buffer

	cmd := exec.Command("python", "checkers", id, solution)
	cmd.Stdout = &out
	cmd.Stderr = &scriptErr

	err := cmd.Run()
	if err != nil {
		return types.Error, errors.WithMessage(err, fmt.Sprintf("python checker failed: %s", scriptErr.String()))
	}

	res, err := strconv.Atoi(out.String())
	if err != nil {
		return types.Error, errors.WithMessage(err, "not int response from python checker")
	}

	return types.ProblemResult(res), nil
}

func InsertSubmit(ctx context.Context, c echo.Context, id int, status types.ProblemResult, solution string) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		c.Logger().Error("assert token failed in post problem by id")

		return
	}

	claims, ok := user.Claims.(*auth.JWTCustomClaims)
	if !ok {
		c.Logger().Error("assert claims failed in post problem by id")

		return
	}

	submit := models.Submit{
		UserID:      claims.UserID,
		ProblemID:   id,
		Status:      int(status),
		SolutionRaw: solution,
	}

	err := submit.InsertG(ctx, boil.Infer())
	if err != nil {
		c.Logger().Error(errors.WithMessage(err, "insert submit failed in post problem by id"))
	}
}
