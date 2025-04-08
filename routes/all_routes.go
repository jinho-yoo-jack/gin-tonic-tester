package routes

import "github.com/gin-gonic/gin"

func AllRouters(
	userRouter func(*gin.Engine),
) []func(*gin.Engine) {
	return []func(*gin.Engine){userRouter}
}
