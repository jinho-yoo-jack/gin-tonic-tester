package server

import (
	"ginTonicProject/config"
	"ginTonicProject/server/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	config config.Config
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(config config.Config, db *gorm.DB) (*Server, error) {
	server := &Server{config: config, db: db}
	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/env", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"env": s.config.Environment})
	})

	routers.UserRouters(router)

	s.router = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
