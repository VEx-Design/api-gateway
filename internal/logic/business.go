package logic

type UserService interface {
	ValidateJWT(tokenString string) (string, error)
}
