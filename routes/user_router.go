package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/handler"
)

type UserRouters struct {
	userHandler *handler.UserHandler
}

func NewUserRouters(userHandler *handler.UserHandler) *UserRouters {
	return &UserRouters{userHandler: userHandler}
}

func (u *UserRouters) RegisterRoutes(router *gin.Engine, authorize gin.HandlerFunc) {
	api := router.Group("/api/v1/auth")
	{
		api.POST("/signup", u.userHandler.SignUpHandler)
		api.POST("/signin", u.userHandler.SignInHandler)
	}
}
