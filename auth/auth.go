package auth

import (
	"crypto/rsa"
	"fmt"
	"geometrics/models"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"time"
)

type JwtCustomClaims struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateAccessToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &JwtCustomClaims{
		ID:   user.ID,
		Name: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		IsAdmin: func() bool {
			if user.Type == 2 {
				return true
			} else {
				return false
			}
		}(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	key, err := GetRSAPrivateKey()
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

func GetRSAPublicKey() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile("./id.rsa.pub.pkcs8")
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

func GetRSAPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile("./id.rsa")
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(keyData)
}
