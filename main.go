//go:generate sqlboiler psql -c sqlboiler.toml
// ssh-keygen -t rsa -m PEM
// ssh-keygen -f id_rsa.pub -e -m pkcs8 > id_rsa.pub.pkcs8

package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/boil"

	adminHandlers "geometrics/handlers/admin"
	APIHandlers "geometrics/handlers/api"
	coursesHandlers "geometrics/handlers/courses"
	problemsHandlers "geometrics/handlers/problems"
	profilesHandlers "geometrics/handlers/profiles"
	"geometrics/types"
	"geometrics/utils"
)

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
	e.HTTPErrorHandler = utils.CustomHTTPErrorHandler
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Accessible")
	})
	e.File("/login", "public/login.html")

	restricted := e.Group("")
	{
		utils.UseAuth(restricted)

		restricted.File("test", "public/test.html")

		restricted.GET("/profiles", profilesHandlers.GETProfile)
		restricted.GET("/profiles/:id", profilesHandlers.GETProfileByID(ctx))

		restricted.GET("/problems/:id", problemsHandlers.GETProblemByID(ctx))

		restricted.GET("/courses/:id", coursesHandlers.GETCourseByID(ctx))
	}

	// API group
	apiG := e.Group("/api/v1")
	{
		apiG.File("", "public/indexAPI.html")
		apiG.POST("/login", APIHandlers.Login(ctx))

		// Problems group
		problems := apiG.Group("/problems")
		{
			utils.UseAuth(problems)

			problems.GET("/:id", APIHandlers.GETProblemByID)
			problems.POST("/:id", APIHandlers.POSTProblemByID)
			problems.PUT("/:id", APIHandlers.PUTProblemByID)
			problems.DELETE("/:id", APIHandlers.DELETEProblemByID)
		}
	}

	// Admin group
	admin := e.Group("/admin")
	{
		utils.UseAuth(admin)

		admin.GET("", adminHandlers.Admin)
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
