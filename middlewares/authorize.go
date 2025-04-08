package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinho-yoo-jack/gin-tonic-tester/config"
	"github.com/jinho-yoo-jack/gin-tonic-tester/model"
	"net/http"
	"strings"
)

func NewAuthorizeMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Fail(401, "Authorization Header missing"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Fail(401, "Invalid Authorization Header"))
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return cfg.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Fail(401, "Invalid Token"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userId, ok := claims["userId"]; ok {
				c.Set("userId", userId)
				c.Next() // Execute Next Filter
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Fail(401, "Unauthorized"))
	}

}
