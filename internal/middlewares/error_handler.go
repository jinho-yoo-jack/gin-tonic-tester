package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/model"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			fmt.Println("Execute ErrorMiddleware")
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(
					http.StatusUnauthorized, model.Fail(http.StatusInternalServerError, err.(error).Error()))
			}
		}()
		c.Next()
	}
}
