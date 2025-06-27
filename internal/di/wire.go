//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jinho-yoo-jack/gin-tonic-tester/config"
	"github.com/jinho-yoo-jack/gin-tonic-tester/handler"
	"github.com/jinho-yoo-jack/gin-tonic-tester/internal/middlewares"
	"github.com/jinho-yoo-jack/gin-tonic-tester/internal/utils"
	repository2 "github.com/jinho-yoo-jack/gin-tonic-tester/persistence/minio"
	"github.com/jinho-yoo-jack/gin-tonic-tester/persistence/minio/repository"
	"github.com/jinho-yoo-jack/gin-tonic-tester/persistence/mysql"
	"github.com/jinho-yoo-jack/gin-tonic-tester/routes"
	"github.com/jinho-yoo-jack/gin-tonic-tester/server"
	"github.com/jinho-yoo-jack/gin-tonic-tester/service"
)

var ConfigSet = wire.NewSet(
	config.MustLoadConfig,
	config.InitDB,
	repository2.NewMinio,
	utils.NewJwtUtils,
)

var MiddlewareSet = wire.NewSet(
	middlewares.NewAuthorizeMiddleware,
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		ConfigSet,
		MiddlewareSet,
		mysql.NewUserRepository,
		repository.NewMinioRepository,
		service.NewUserService,
		service.NewMinioService,
		handler.NewUserHandler,
		handler.NewMinioHandler,
		routes.NewUserRouters,
		routes.NewMinioRouters,
		routes.NewAllOfRouters,
		//wire.Value([]func(*gin.Engine){
		//	routes.NewUserRouters,
		//}),
		server.NewServer,
	)
	return nil, nil
}
