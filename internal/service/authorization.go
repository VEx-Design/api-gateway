package service

import (
	"errors"
	"os"
	"time"

	"github.com/Atipat-CMU/api-gateway/internal/logic"
	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
}

func NewAuthService() logic.AuthService {
	return &authService{}
}

func (u *authService) ValidateJWT(tokenString string) (string, error) {

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method matches the expected one
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return "", err // Return the error if parsing fails
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the user ID from the claims
		userId, ok := claims["id"].(string)
		if !ok {
			return "", errors.New("user ID not found in token")
		}

		// Optionally, check if the token is expired
		expiration, ok := claims["exp"].(float64)
		if ok && time.Now().Unix() > int64(expiration) {
			return "", errors.New("token expired")
		}

		return userId, nil
	}

	return "", errors.New("invalid token")
}
