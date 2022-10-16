//go:generate sqlboiler psql -c sqlboiler.toml
// ssh-keygen -t rsa -m PEM
// ssh-keygen -f id.rsa.pub -e -m pkcs8 > id.rsa.pub.pkcs8

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
	"github.com/volatiletech/sqlboiler/v4/boil"

	adminHandlers "geometrics/handlers/admin"
	APIHandlers "geometrics/handlers/api"
	coursesHandlers "geometrics/handlers/courses"
	problemsHandlers "geometrics/handlers/problems"
	profilesHandlers "geometrics/handlers/profiles"
	submitsHandlers "geometrics/handlers/submits"
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
	e.Renderer = &types.Template{Templates: template.Must(template.New("").Funcs(
		template.FuncMap{
			"statusStringify": func(status int) string { return types.ProblemResult(status).String() },
		},
	).ParseGlob("public/*.gohtml"))}
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
	e.Use(utils.LoggerMiddleware())

	// Handlers
	e.File("", "public/index.gohtml")

	e.File("/login", "public/login.html")

	e.GET("/profiles", profilesHandlers.GETProfile, utils.AuthMiddleware(true))
	e.GET("/profiles/:id", profilesHandlers.GETProfileByID(ctx))

	e.GET("/problems/:id", problemsHandlers.GETProblemByID(ctx), utils.AuthMiddleware(true))
	e.GET("/problems/:id/submits", problemsHandlers.GETSubmitsByID(ctx), utils.AuthMiddleware(true))
	e.GET("/problems/:id/solve", problemsHandlers.GETSolveByID(ctx), utils.AuthMiddleware(false))

	e.GET("/submits/:id", submitsHandlers.GETSubmitsByID(ctx), utils.AuthMiddleware(false))

	e.GET("/courses/:id", coursesHandlers.GETCourseByID(ctx))

	// API group
	apiG := e.Group("/api/v1")
	{
		apiG.File("", "public/indexAPI.html")
		apiG.POST("/login", APIHandlers.Login(ctx))

		// Problems group
		problems := apiG.Group("/problems")
		{
			problems.Use(utils.AuthMiddleware(false))

			problems.GET("/:id", APIHandlers.GETProblemByID)
			problems.POST("/:id", APIHandlers.POSTProblemByID(ctx))
			problems.PUT("", APIHandlers.PUTProblem(ctx))
			problems.PATCH("/:id", APIHandlers.PATCHProblemByID(ctx))
			problems.DELETE("/:id", APIHandlers.DELETEProblemByID(ctx))
		}
	}

	// Admin group
	admin := e.Group("/admin")
	{
		admin.Use(utils.AuthMiddleware(false))

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

	ctx, cancel := context.WithTimeout(ctx, time.Second*10) //nolint:gomnd
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
