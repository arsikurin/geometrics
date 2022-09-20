//go:generate sqlboiler psql -c sqlboiler.toml
// ssh-keygen -t rsa -m PEM
// ssh-keygen -f id_rsa.pub -e -m pkcs8 > id_rsa.pub.pkcs8

package main

import (
	"context"
	"database/sql"
	"fmt"
	"geometrics/auth"
	apiH "geometrics/handlers/api"
	"geometrics/types"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Accessible")
	}
}

func restricted(c echo.Context) error {
	//cookie, err := c.Cookie("token")
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		return c.String(http.StatusUnauthorized, "no cookie")
	//	}
	//	return c.String(http.StatusBadRequest, err.Error())
	//}
	//
	//claims := &types.JwtCustomClaims{}
	//
	//tkn, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
	//	return auth.GetRSAPublicKey()
	//})
	//if err != nil {
	//	if err == jwt.ErrSignatureInvalid {
	//		return c.String(http.StatusUnauthorized, "sig invalid")
	//	}
	//	return c.String(http.StatusBadRequest, err.Error())
	//}
	//if !tkn.Valid {
	//	return c.String(http.StatusUnauthorized, "token invalid")
	//}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtCustomClaims)

	return c.String(http.StatusOK, "Welcome "+claims.Name+"!")
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	title := fmt.Sprintf("%d Internal Server Error", code)
	var detail []types.ApiError

	if he, ok := err.(*echo.HTTPError); ok {
		if detail, ok = he.Message.([]types.ApiError); ok {
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
			//"detail": detail,
		}); err != nil {
			c.Logger().Error(err)
		}
	}
	c.Logger().Error(err)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	ctx := context.Background()
	db, err := sql.Open("postgres", fmt.Sprintf(
		"dbname=%s host=%s user=%s password=%s sslmode=require",
		os.Getenv("DBNAME"), os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"),
	))
	if err != nil {
		e.Logger.Error(err)
	}
	boil.SetDB(db)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			e.Logger.Fatal(err.Error())
		}
	}(db)

	e.Debug = true
	e.Renderer = &types.Template{Templates: template.Must(template.ParseGlob("public/*.html"))}
	e.Validator = &types.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Static("/static", "public/static")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	e.GET("/", index())
	e.File("/login", "public/login.html")

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
				Claims:        &auth.JwtCustomClaims{},
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

	// Restricted group
	admin := e.Group("/admin")
	{
		key, err := auth.GetRSAPublicKey()
		if err != nil {
			log.Println(err)
		}

		admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:        &auth.JwtCustomClaims{},
			SigningKey:    key,
			TokenLookup:   "header:Authorization,cookie:token",
			SigningMethod: "RS256",
		}))

		admin.GET("", restricted)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
