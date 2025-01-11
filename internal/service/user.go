package service

import (
	ports "github.com/Atipat-CMU/api-gateway/external/_ports"
	"github.com/Atipat-CMU/api-gateway/internal/entities"
	"github.com/Atipat-CMU/api-gateway/internal/logic"
)

type userService struct {
	userRev ports.UserReceiver
}

func NewUserService(userRev ports.UserReceiver) logic.UserService {
	return &userService{
		userRev: userRev,
	}
}

func (u *userService) GetUser(userId string) (*entities.User, error) {
	return u.userRev.GetUser(userId)
}
