package controller

import (
	"auth/app/document/model"
	"auth/app/document/service"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	service service.IService
}

func NewController(service service.IService) controller {
	return controller{service: service}
}

func (ct controller) UpdateDocument(ctx *fiber.Ctx) error {
	req := model.Request{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set RefId in context
	result := ct.service.UpdateDocument(ctx, &req)
	result.RefID = req.CallbackRefid
	return ctx.JSON(result)

}
