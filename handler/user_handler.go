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

func SignUp(c *gin.Context) {
	var req RequestSignUp

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": service.SignUp(
		service.UserInfo{UserId: req.UserId, Password: req.Password, NickName: req.NickName, Role: 1})},
	)
}
