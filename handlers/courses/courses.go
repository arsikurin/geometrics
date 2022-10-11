package courses

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"geometrics/models"
)

func GETCourseByID(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.ErrNotFound
		}

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
			OrderBy(models.ProblemColumns.Name),
		).AllG(ctx)
		if err != nil {
			return errors.WithMessage(err, "inner join failed in get course by id")
		}

		course, err := models.Courses(
			Select(models.CourseColumns.Name, models.CourseColumns.Description, models.CourseColumns.AuthorID),
			models.CourseWhere.ID.EQ(id),
		).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get course failed in get course by id")
		}

		author, err := models.Users(
			Select(models.UserColumns.FirstName, models.UserColumns.LastName),
			models.UserWhere.ID.EQ(course.AuthorID),
		).OneG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get author failed in get course by id")
		}

		return c.Render(http.StatusOK, "course.gohtml", map[string]interface{}{
			"author":   fmt.Sprintf("%s %s", author.FirstName, author.LastName),
			"course":   course,
			"problems": courseProblems,
		})
	}
}
