// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/adapter/controller"
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/adapter/gateway"
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/infrastructure/bun"
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/usecase"
)

// Injectors from wire.go:

func Init() *controller.UserHandler {
	db := bun.NewDB()
	userRepository := gateway.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := controller.NewUserHandler(db, userUsecase)
	return userHandler
}