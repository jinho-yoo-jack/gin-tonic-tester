//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/jinho-yoo-jack/gin-tonic-tester/config"
	"github.com/jinho-yoo-jack/gin-tonic-tester/handler"
	"github.com/jinho-yoo-jack/gin-tonic-tester/middlewares"
	"github.com/jinho-yoo-jack/gin-tonic-tester/repository"
	"github.com/jinho-yoo-jack/gin-tonic-tester/routes"
	"github.com/jinho-yoo-jack/gin-tonic-tester/server"
	"github.com/jinho-yoo-jack/gin-tonic-tester/service"
	"github.com/jinho-yoo-jack/gin-tonic-tester/utils"
)

var ConfigSet = wire.NewSet(
	config.MustLoadConfig,
	config.InitDB,
	utils.NewJwtUtils,
)

var MiddlewareSet = wire.NewSet(
	middlewares.NewAuthorizeMiddleware,
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		ConfigSet,
		MiddlewareSet,
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
		routes.NewUserRouters,
		routes.AllRouters,
		//wire.Value([]func(*gin.Engine){
		//	routes.NewUserRouters,
		//}),
		server.NewServer,
	)
	return nil, nil
}
