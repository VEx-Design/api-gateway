package handler

import (
	"net/http"

	"github.com/Atipat-CMU/api-gateway/internal/logic"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSrv logic.UserService
}

func NewUserHandler(userSrv logic.UserService) *UserHandler {
	return &UserHandler{
		UserSrv: userSrv,
	}
}

func (h *UserHandler) GetID(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")
	if userId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing X-User-Id header"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userId": userId})
}
