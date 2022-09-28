//go:generate sqlboiler psql -c sqlboiler.toml
// ssh-keygen -t rsa -m PEM
// ssh-keygen -f id_rsa.pub -e -m pkcs8 > id_rsa.pub.pkcs8

package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"geometrics/auth"
	apiH "geometrics/handlers/api"
	"geometrics/models"
	"geometrics/types"
)

func index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Accessible")
	}
}

func adminPage(c echo.Context) error {
	// cookie, err := c.Cookie("token")
	// if err != nil {
	//	if err == http.ErrNoCookie {
	//		return c.String(http.StatusUnauthorized, "no cookie")
	//	}
	//	return c.String(http.StatusBadRequest, err.Error())
	// }
	//
	// claims := &types.JWTCustomClaims{}
	//
	// tkn, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
	//	return auth.GetRSAPublicKey()
	// })
	// if err != nil {
	//	if err == jwt.ErrSignatureInvalid {
	//		return c.String(http.StatusUnauthorized, "sig invalid")
	//	}
	//	return c.String(http.StatusBadRequest, err.Error())
	// }
	// if !tkn.Valid {
	//	return c.String(http.StatusUnauthorized, "token invalid")
	// }
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTCustomClaims)

	return c.String(http.StatusOK, "["+strconv.FormatBool(claims.IsAdmin)+"] Welcome "+claims.Name+"!")
}

func customHTTPErrorHandler(err error, c echo.Context) {
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

func main() {
	fmt.Println(`
   ______                               __         _            
  / ____/___   ____   ____ ___   ___   / /_ _____ (_)_____ _____
 / / __ / _ \ / __ \ / __ '__ \ / _ \ / __// ___// // ___// ___/
/ /_/ //  __// /_/ // / / / / //  __// /_ / /   / // /__ (__  ) 
\____/ \___/ \____//_/ /_/ /_/ \___/ \__//_/   /_/ \___//____/  

	`)

	e := echo.New()
	e.Debug = true
	e.Logger.SetHeader(`{"level":"${level}","time":"${time_rfc3339}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}`)
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)

	if e.Debug {
		err := godotenv.Load(".env.development.local")
		if err != nil {
			e.Logger.Fatal("Error loading .env.development.local file")
		}
	}

	db, err := sql.Open("postgres", fmt.Sprintf(
		"dbname=%s host=%s user=%s password=%s sslmode=require",
		os.Getenv("PG_DBNAME"), os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"),
	))
	if err != nil {
		e.Logger.Error(err)
	}
	boil.SetDB(db)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			e.Logger.Error(err.Error())
		}
	}(db)

	e.HideBanner = true
	e.Renderer = &types.Template{Templates: template.Must(template.ParseGlob("public/*.html"))}
	e.Validator = &types.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Static("/static", "public/static")

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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
	}))

	e.GET("/", index())
	e.File("/login", "public/login.html")

	restricted := e.Group("")
	{
		key, err := auth.GetRSAPublicKey()
		if err != nil {
			log.Println(err)
		}

		restricted.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:        &auth.JWTCustomClaims{},
			SigningKey:    key,
			TokenLookup:   "header:Authorization,cookie:token",
			SigningMethod: "RS256",
		}))

		restricted.File("test", "public/test.html")

		restricted.GET("/profiles", func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*auth.JWTCustomClaims)

			return c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/profiles/%d", claims.ID))
		})
		restricted.GET("/profiles/:id", func(c echo.Context) error {
			id := c.Param("id")
			claims := c.Get("user").(*jwt.Token).Claims.(*auth.JWTCustomClaims)
			name := claims.Name

			if isExists, err := models.Users(Where("id=?", claims.Id)).ExistsG(ctx); !isExists {
				if err != nil {
					return err
				}
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"code":   http.StatusUnauthorized,
					"status": "error",
					"message": fmt.Sprintf(
						"%d %s", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized),
					),
					"detail": "user not exists",
				})
			}
			user, err := models.Users(Where("id=?", claims.ID)).OneG(ctx)
			if err != nil {
				return err
			}

			return c.String(http.StatusOK, fmt.Sprintf("profile %s %s %d %v %v %s %s", id, name, user.Type, user.Grade, user.School, user.CreatedAt, user.Timezone))
		})

		restricted.GET("/problems/:id", func(c echo.Context) error {
			id := c.Param("id")
			action := c.QueryParam("action")
			if action == "" {
				action = "view"
			}
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*auth.JWTCustomClaims)
			name := claims.Name

			return c.String(http.StatusOK, "problem "+id+" "+name)
		})

		restricted.GET("/courses/:id", func(c echo.Context) error {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return errors.WithMessage(err, "Not int ID /courses")
			}
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*auth.JWTCustomClaims)

			if isExists, err := models.Courses(Where("id=?", id)).ExistsG(ctx); !isExists {
				if err != nil {
					return err
				}
				return echo.ErrNotFound
			}

			courseProblems, err := models.Problems(
				InnerJoin(fmt.Sprintf("%s on %s=%s",
					models.TableNames.CoursesProblems, models.CoursesProblemTableColumns.ProblemID, models.ProblemTableColumns.ID)),
				models.CoursesProblemWhere.CourseID.EQ(id),
			).AllG(ctx)
			if err != nil {
				return err
			}

			for _, problem := range courseProblems {
				fmt.Println(problem.ID)
				fmt.Println(problem.Name)
				fmt.Println(problem.Description)
			}

			return c.String(http.StatusOK, "course "+claims.Name)
		})
	}

	// API group
	apiG := e.Group("/api/v1")
	{
		apiG.File("", "public/indexAPI.html")
		apiG.POST("/login", apiH.Login(ctx))

		// Problems group
		problems := apiG.Group("/problems")
		{
			key, err := auth.GetRSAPublicKey()
			if err != nil {
				log.Println(err)
			}

			problems.Use(middleware.JWTWithConfig(middleware.JWTConfig{
				Claims:        &auth.JWTCustomClaims{},
				SigningKey:    key,
				TokenLookup:   "header:Authorization,cookie:token",
				SigningMethod: "RS256",
			}))

			problems.GET("/:id", apiH.ProblemsGET)
			problems.POST("/:id", apiH.ProblemsPOST)
			problems.PUT("/:id", apiH.ProblemsPUT)
			problems.DELETE("/:id", apiH.ProblemsDELETE)
		}
	}

	// Admin group
	admin := e.Group("/admin")
	{
		key, err := auth.GetRSAPublicKey()
		if err != nil {
			log.Println(err)
		}

		admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:        &auth.JWTCustomClaims{},
			SigningKey:    key,
			TokenLookup:   "header:Authorization,cookie:token",
			SigningMethod: "RS256",
		}))

		admin.GET("", adminPage)
	}

	// Graceful shutdown
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
