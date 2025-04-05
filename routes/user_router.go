package routes

import (
	"ginTonicProject/injector"
	"github.com/gin-gonic/gin"
)

func UserRouters(router *gin.Engine) {
	userHandler := injector.InitializeUserHandler()
	api := router.Group("/api/v1/user")
	{
		api.POST("/signup", userHandler.SignUpHandler)
	}
}
