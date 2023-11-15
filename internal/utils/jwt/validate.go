package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"

	"github.com/Shemistan/uzum_auth/internal/models"
)

func ValidateToken(tokenString, secretKey string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
