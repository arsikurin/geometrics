package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"geometrics/models"
	"geometrics/types"
)

type JWTCustomClaims struct {
	UserID  int    `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"` //nolint:tagliatelle
	jwt.StandardClaims
}

func GenerateAccessToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &JWTCustomClaims{
		UserID:  user.ID,
		Name:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		IsAdmin: user.Type == int(types.Admin),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //nolint:gomnd
		},
	})

	key, err := GetRSAPrivateKey()
	if err != nil {
		return "", err
	}

	return token.SignedString(key) //nolint:wrapcheck
}

func GetRSAPublicKey() (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile("./id.rsa.pub.pkcs8")
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return jwt.ParseRSAPublicKeyFromPEM(keyData) //nolint:wrapcheck
}

func GetRSAPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile("./id.rsa")
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return jwt.ParseRSAPrivateKeyFromPEM(keyData) //nolint:wrapcheck
}
