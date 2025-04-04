package handler

import (
	"ginTonicProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestSignUp struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
}

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUpHandler(c *gin.Context) {
	var req RequestSignUp

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": h.userService.SignUp(
		service.UserInfo{UserId: req.UserId, Password: req.Password, NickName: req.NickName, Role: 1})},
	)
}
