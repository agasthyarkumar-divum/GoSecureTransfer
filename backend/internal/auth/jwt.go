package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ⚠️ Temporary secret (we will move this to env later)
var secret = []byte("supersecretkey")

// Generate JWT token
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// Validate JWT token
func ValidateToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	user, ok := claims["user"].(string)
	if !ok {
		return "", err
	}

	return user, nil
}