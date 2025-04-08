package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"mark99/config"
	"mark99/routes"
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
	router := gin.Default()
	routes.HealthcheckRouters(router)
	//routes.UserRouters(router, s.authorize)
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
