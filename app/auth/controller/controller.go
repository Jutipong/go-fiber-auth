package controller

import (
	"auth/app/auth/model"
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
	// Validation Model
	var authReq model.Auth
	if err := authReq.Validation(ctx); err != nil {
		return fiber.ErrBadRequest
	}

	result, err := ct.service.Login(ctx, &authReq)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}
	return ctx.JSON(result)
}

func (ct controller) Create(ctx *fiber.Ctx) error {
	auth := model.Auth{}
	if err := ctx.BodyParser(&auth); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err := ct.service.Create(ctx, &auth)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}
	return ctx.JSON(fiber.Map{"status": "success"})
}
