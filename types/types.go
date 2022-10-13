package types

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ProblemResult int

const (
	OK      ProblemResult = iota
	WA      ProblemResult = iota
	Invalid ProblemResult = iota
	Error   ProblemResult = iota
)

func (pr ProblemResult) String() string {
	switch pr {
	case OK:
		return "Correct"
	case WA:
		return "Wrong answer"
	case Invalid:
		return "Forbidden"
	case Error:
		return "Error occurred"
	default:
		return "unknown"
	}
}

type UserType int

const (
	Student UserType = iota
	Teacher UserType = iota
	Admin   UserType = iota
)

func (ut UserType) String() string {
	switch ut {
	case Student:
		return "Student"
	case Teacher:
		return "Teacher"
	case Admin:
		return "Admin"
	default:
		return "unknown"
	}
}

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data) //nolint:wrapcheck
}

type (
	POSTProblemReq struct {
		GgbBase64 string `json:"ggb_base64" validate:"required,base64"` //nolint:tagliatelle
	}

	PUTProblemReq struct {
		Name           string `json:"name" validate:"required"`
		Description    string `json:"description" validate:"required"`
		SolutionBase64 string `json:"solution_base64" validate:"required,base64"` //nolint:tagliatelle
	}

	PATCHProblemReq struct {
		Name           string `json:"name,omitempty"`
		Description    string `json:"description,omitempty"`
		SolutionBase64 string `json:"solution_base64,omitempty" validate:"omitempty,base64"` //nolint:tagliatelle
	}

	LoginReq struct {
		Login    string `json:"login" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,lte=30"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}

	APIError struct {
		Field   string
		Message string
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		if ves, ok := err.(validator.ValidationErrors); ok {
			out := make([]APIError, len(ves))

			for i, ve := range ves {
				switch ve.Tag() {
				case "required":
					out[i] = APIError{
						Field:   ve.Field(),
						Message: "This field is required",
					}
				case "email":
					out[i] = APIError{
						Field:   ve.Field(),
						Message: "Invalid email",
					}
				case "gte":
					out[i] = APIError{
						Field:   ve.Field(),
						Message: fmt.Sprintf("Must be greater than %s", ve.Param()),
					}
				case "lte":
					out[i] = APIError{
						Field:   ve.Field(),
						Message: fmt.Sprintf("Must be less than %s", ve.Param()),
					}
				default:
					out[i] = APIError{
						Field:   ve.Field(),
						Message: err.Error(),
					}
				}
			}

			return echo.NewHTTPError(http.StatusBadRequest, out)
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
