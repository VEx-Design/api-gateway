package logic

import "github.com/Atipat-CMU/api-gateway/internal/entities"

type AuthService interface {
	ValidateJWT(tokenString string) (string, error)
}

type UserService interface {
	GetUser(userId string) (*entities.User, error)
}
