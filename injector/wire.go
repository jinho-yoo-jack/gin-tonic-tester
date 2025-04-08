//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"mark99/config"
	"mark99/handler"
	"mark99/middlewares"
	"mark99/repository"
	"mark99/routes"
	"mark99/server"
	"mark99/service"
	"mark99/utils"
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
