package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/model"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, model.Fail(http.StatusInternalServerError, "Server Error"))
			}
		}()
		c.Next()
	}
}
