package query

import (
	"github.com/Atipat-CMU/api-gateway/external/handler/adaptors/graphql/model"
	"github.com/Atipat-CMU/api-gateway/internal/logic"
)

type UserQuery interface {
	GetUser(userId string) (*model.User, error)
}

type userQuery struct {
	userSrv logic.UserService
}

func NewUserQuery(userSrv logic.UserService) UserQuery {
	return &userQuery{userSrv: userSrv}
}

func (q *userQuery) GetUser(userId string) (*model.User, error) {
	user, err := q.userSrv.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
	}, nil
}
