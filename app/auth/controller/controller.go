package controller

import (
	"auth/app/auth/service"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	service service.IService
}

func NewController(service service.IService) controller {
	return controller{service: service}
}

func (ct controller) Login(ctx *fiber.Ctx) error {
	result, err := ct.service.Login(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}
	return ctx.JSON(result)
}
