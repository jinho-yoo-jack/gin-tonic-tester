package routers

import (
	"ginTonicProject/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userHandler *handler.UserHandler
}

func (r *UserRouter) UserRouters(router *gin.Engine) {
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/signup", r.userHandler.SignUpHandler)
	}
}
