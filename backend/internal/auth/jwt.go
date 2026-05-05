package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// secret is initialized from environment or default
var secret []byte

// InitAuth initializes the auth package with environment variables
func InitAuth() {
	// Read JWT secret from environment
	secretStr := os.Getenv("JWT_SECRET")
	if secretStr == "" {
		secretStr = "supersecretkey" // Default secret
		log.Println("⚠️ JWT_SECRET not set, using default")
	}

	if len(secretStr) < 32 {
		log.Println("⚠️ JWT_SECRET should be at least 32 characters for security")
	}

	secret = []byte(secretStr)
}

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