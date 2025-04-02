package routers

import (
	"ginTonicProject/handler"
	"github.com/gin-gonic/gin"
)

func UserRouters(router *gin.Engine) {
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/signup", handler.SignUpHandler)
	}
}
