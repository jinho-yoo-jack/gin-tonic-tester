package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthcheckRouters(router *gin.Engine) {
	router.POST("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/env", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"env": "config"})
	})
}
