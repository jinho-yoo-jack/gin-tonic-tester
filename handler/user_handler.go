package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/service"
	"net/http"
)

type RequestSignUp struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
}

type RequestSignIn struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
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

	c.Set("handler_result", h.userService.SignUp(
		&service.UserInfo{UserId: req.UserId, Password: req.Password, NickName: req.NickName, Role: 1}))

	//c.JSON(http.StatusOK, gin.H{"user":},
}

func (h *UserHandler) SignInHandler(c *gin.Context) {
	var req RequestSignIn

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": h.userService.SignIn(service.SignInDto{UserId: req.UserId, Password: req.Password})})

}
