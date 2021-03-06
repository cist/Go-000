// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"Week04/internal/biz"
	"Week04/internal/data"
	"Week04/internal/service"
)

// Injectors from wire.go:

func InitUserService() *service.UserService {
	userRepo := data.NewUserRepo()
	userUsecase := biz.NewUserUsecase(userRepo)
	userService := service.NewUserService(userUsecase)
	return userService
}
