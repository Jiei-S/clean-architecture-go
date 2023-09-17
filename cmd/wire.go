//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/adapter/controller"
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/adapter/gateway"
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/infrastructure/bun"
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/usecase"

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
