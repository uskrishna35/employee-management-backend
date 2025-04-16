package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing JWT (store this securely)
var secretKey = []byte("your_secret_key")

// GenerateJWT creates a JWT token with an email claim and expiration time
func GenerateJWT(email string) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	}

	// Create new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(secretKey)
}
