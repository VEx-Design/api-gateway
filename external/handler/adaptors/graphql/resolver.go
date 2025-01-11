package graph

import "github.com/Atipat-CMU/api-gateway/internal/logic"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserSrv logic.UserService
}
