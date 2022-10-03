package courses

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"geometrics/auth"
	"geometrics/models"
)

func GETCourseByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.JWTCustomClaims)

		if isExists, err := models.Courses(Where("id=?", id)).ExistsG(ctx); !isExists {
			if err != nil {
				return errors.WithMessage(err, "check whether user exists failed in get course by id")
			}

			return echo.ErrNotFound
		}

		courseProblems, err := models.Problems(
			InnerJoin(fmt.Sprintf("%s on %s=%s",
				models.TableNames.CoursesProblems, models.CoursesProblemTableColumns.ProblemID, models.ProblemTableColumns.ID)),
			models.CoursesProblemWhere.CourseID.EQ(id),
		).AllG(ctx)
		if err != nil {
			return errors.WithMessage(err, "inner join failed in get course by id")
		}

		for _, problem := range courseProblems {
			fmt.Println(problem.ID)
			fmt.Println(problem.Name)
			fmt.Println(problem.Description)
		}

		return c.String(http.StatusOK, "course "+claims.Name)
	}
}
