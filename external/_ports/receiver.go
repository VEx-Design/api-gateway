package ports

import "github.com/Atipat-CMU/api-gateway/internal/entities"

type UserReceiver interface {
	GetUser(userId string) (*entities.User, error)
}
