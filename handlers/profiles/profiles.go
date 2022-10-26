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
	"geometrics/types"
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
			return errors.WithMessage(err, "get user failed in get profile by id")
		}

		solvedProblems, err := models.Submits(
			Where("user_id=?", id),
			Where("status=?", int(types.OK)),
			Distinct("problem_id"),
		).AllG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get solved problems failed in get profile by id")
		}

		return c.Render(http.StatusOK, "profile.gohtml", map[string]interface{}{ //nolint:wrapcheck
			"name":           fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			"user":           user,
			"solvedProblems": len(solvedProblems),
		})
	}
}

func GETProfile(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("assert token failed in get profile")
	}

	claims, ok := user.Claims.(*auth.JWTCustomClaims)
	if !ok {
		return errors.New("assert claims failed in get profile")
	}
	fmt.Println(claims.UserID)
	fmt.Println(claims.Name)

	if claims.UserID != -1 {
		return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/profiles/%d", claims.UserID)) //nolint:wrapcheck
	}

	return c.Redirect(http.StatusPermanentRedirect, "/login") //nolint:wrapcheck
}
