package routes

import (
	"auth/app/auth/controller"
	"auth/app/auth/repository"
	"auth/app/auth/service"
	"auth/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(ctx *fiber.App) {
	repository := repository.NewRepository(config.Db())
	service := service.NewService(repository)
	controller := controller.NewController(service)
	ctx.Post("/UpdateDocument", controller.UpdateDocument)
}
