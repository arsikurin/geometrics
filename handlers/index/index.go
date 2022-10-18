package index

import (
	"context"
	"net/http"

	"github.com/friendsofgo/errors"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/labstack/echo/v4"

	"geometrics/models"
)

func GETIndex(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		courses, err := models.Courses(
			Load("CoursesProblems"),
		).AllG(ctx)
		if err != nil {
			return errors.WithMessage(err, "get courses failed in get index")
		}

		return c.Render(http.StatusOK, "index.gohtml", map[string]interface{}{ //nolint:wrapcheck
			"courses": courses,
			"ctx":     ctx,
		})
	}
}

// author_name (first_name + " " + last_name) + array of problems for each course
