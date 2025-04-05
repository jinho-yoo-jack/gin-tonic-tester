package server

import (
	"ginTonicProject/config"
	"ginTonicProject/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(config *config.Config, db *gorm.DB) (*Server, error) {
	server := &Server{config: config, db: db}
	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	routes.HealthcheckRouters(router)
	routes.UserRouters(router)

	s.router = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
