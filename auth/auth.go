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
	IsAdmin bool   `json:"is_admin"` // is admin json check register
	jwt.StandardClaims
}

func GenerateAccessToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &JWTCustomClaims{
		UserID: user.ID,
		Name:   fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		IsAdmin: func() bool {
			if user.Type == int(types.Admin) {
				return true
			}
			return false
		}(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //nolint:gomnd
		},
	})

	key, err := GetRSAPrivateKey()
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

func GetRSAPublicKey() (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile("./id.rsa.pub.pkcs8")
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

func GetRSAPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile("./id.rsa")
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPrivateKeyFromPEM(keyData)
}
