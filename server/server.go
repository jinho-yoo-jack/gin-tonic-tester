package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinho-yoo-jack/gin-tonic-tester/config"
	"github.com/jinho-yoo-jack/gin-tonic-tester/internal/middlewares"
	"github.com/jinho-yoo-jack/gin-tonic-tester/routes"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(cfg *config.Config, db *gorm.DB, allRouters []func(*gin.Engine)) (*Server, error) {
	server := &Server{
		config: cfg,
		db:     db,
	}
	server.setupRouter(allRouters)
	return server, nil
}

func (s *Server) setupRouter(routers []func(*gin.Engine)) {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middlewares.ErrorMiddleware())
	routes.HealthcheckRouters(router)
	for _, fn := range routers {
		fn(router)
	}

	s.router = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) {
	//return s.router.Run(address)
	err := s.router.Run(address)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
