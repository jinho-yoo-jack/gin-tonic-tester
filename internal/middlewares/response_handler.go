package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/model"
	"net/http"
)

func SuccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			fmt.Println("Success Response Handler")
			result, isExist := c.Get("handler_result")
			if !isExist {
				panic("Not Have Response Message")
			}
			c.JSON(http.StatusOK, model.Success(result))
		}()
		c.Next()
	}
}

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
