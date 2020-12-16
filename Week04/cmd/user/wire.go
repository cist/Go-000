//+build wireinject
// to create InitUserService

package main

import (
	"Week04/internal/biz"
	"Week04/internal/data"
	"Week04/internal/service"
	"github.com/google/wire"
)

func InitUserService() *service.UserService {
	wire.Build(service.NewUserService, biz.NewUserUsecase, data.NewUserRepo)
	return &service.UserService{}
}
