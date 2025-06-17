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
	config      *config.Config
	db          *gorm.DB
	engin       *gin.Engine
	middlewares gin.HandlerFunc
	routers     *routes.AllRouters
}

func NewServer(cfg *config.Config, db *gorm.DB, middleware gin.HandlerFunc, allRouters *routes.AllRouters) (*Server, error) {
	server := &Server{
		config:      cfg,
		db:          db,
		middlewares: middleware,
		routers:     allRouters,
	}
	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middlewares.ErrorMiddleware())
	router.Use(middlewares.SuccessMiddleware())
	routes.HealthcheckRouters(router)
	s.routers.RegisterAllRoutes(router, s.middlewares)

	s.engin = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) {
	//return s.router.Run(address)
	err := s.engin.Run(address)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
