package admin

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"geometrics/auth"
)

func Admin(c echo.Context) error {
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
