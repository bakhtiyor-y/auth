package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/Shemistan/uzum_auth/internal/models"
)

const (
	accessTTL  = time.Hour * 24 * 30
	refreshTTL = time.Hour * 24 * 365
)

func GenerateTokens(login, role, secretKey string) (models.Token, error) {
	var res models.Token

	// Генерация Access Token
	accessClaims := models.CustomClaims{
		Login: login,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTTL).Unix(), // Срок действия токена
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err.Error())
		return res, err
	}

	// Генерация Refresh Token
	refreshClaims := models.CustomClaims{
		Login: login,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTTL).Unix(), // Срок действия токена
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err.Error())
		return res, err
	}

	res.Access = accessString
	res.Refresh = refreshString

	return res, nil
}
