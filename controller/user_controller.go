package controller

import (
	"ginTonicProject/model"
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

	user := model.User{}
	user.UserId = req.UserId
	user.Password = req.Password
	user.NickName = req.NickName
	user.Role = 1

	savedUser, err := user.Save()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"user": savedUser})

}
