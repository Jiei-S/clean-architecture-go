//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/adapter/controller"
	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/adapter/gateway"
	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/infrastructure/bun"
	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/usecase"

	"github.com/google/wire"
)

func Init() *controller.UserHandler {
	wire.Build(
		controller.NewUserHandler,
		bun.NewDB,
		usecase.NewUserUsecase,
		gateway.NewUserRepository,
	)
	return &controller.UserHandler{}
}
