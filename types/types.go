package types

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

type (
	LoginCreds struct {
		Login    string `json:"login" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,lte=30"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}

	ApiError struct {
		Field   string
		Message string
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		if ves, ok := err.(validator.ValidationErrors); ok {
			out := make([]ApiError, len(ves))

			for i, ve := range ves {
				switch ve.Tag() {
				case "required":
					out[i] = ApiError{
						Field:   ve.Field(),
						Message: "This field is required",
					}
				case "email":
					out[i] = ApiError{
						Field:   ve.Field(),
						Message: "Invalid email",
					}
				case "gte":
					out[i] = ApiError{
						Field:   ve.Field(),
						Message: fmt.Sprintf("Must be greater than %s", ve.Param()),
					}
				case "lte":
					out[i] = ApiError{
						Field:   ve.Field(),
						Message: fmt.Sprintf("Must be less than %s", ve.Param()),
					}
				default:
					out[i] = ApiError{
						Field:   ve.Field(),
						Message: err.Error(),
					}
				}
			}
			return echo.NewHTTPError(http.StatusBadRequest, out)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}
	return nil
}
