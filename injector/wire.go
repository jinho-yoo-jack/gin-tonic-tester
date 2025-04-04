//go:build wireinject
// +build wireinject

package injector

import (
	"ginTonicProject/handler"
	"ginTonicProject/service"
	"github.com/google/wire"
)

func InitializeUserHandler() *handler.UserHandler {
	wire.Build(service.NewUserService, handler.NewUserHandler)
	return nil
}
