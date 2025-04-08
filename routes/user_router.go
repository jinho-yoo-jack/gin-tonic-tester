package routes

import (
	"github.com/gin-gonic/gin"
	"mark99/handler"
)

//func UserRouters(router *gin.Engine, authorize gin.HandlerFunc) {
//	userHandler := injector.InitializeUserHandler()
//	api := router.Group("/api/v1/auth")
//	api.Use(authorize)
//	{
//		api.POST("/signup", userHandler.SignUpHandler)
//		api.POST("/signin", userHandler.SignInHandler)
//	}
//}

func NewUserRouters(authorize gin.HandlerFunc, userHandler *handler.UserHandler) func(*gin.Engine) {
	return func(router *gin.Engine) {
		api := router.Group("/api/v1/auth")
		api.Use(authorize)
		{
			api.POST("/signup", userHandler.SignUpHandler)
			api.POST("/signin", userHandler.SignInHandler)
		}
	}
}
