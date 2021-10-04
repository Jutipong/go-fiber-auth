package routes

import (
	"auth/app/document/controller"
	"auth/app/document/repository"
	"auth/app/document/service"
	"auth/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(ctx *fiber.App) {
	repository := repository.NewRepository(config.Db())
	service := service.NewService(repository)
	controller := controller.NewController(service)
	ctx.Post("/UpdateDocument", controller.UpdateDocument)
}
