//go:build wireinject
// +build wireinject

package injector

import (
	"ginTonicProject/config"
	"ginTonicProject/handler"
	"ginTonicProject/repository"
	"ginTonicProject/service"
	"github.com/google/wire"
)

func InitializeUserHandler() *handler.UserHandler {
	wire.Build(
		config.MustLoadConfig,
		config.InitDB,
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler)
	return nil
}
