package utils

import (
	"auth/pkg/enum"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserInfo struct {
	TransactionId string `json:"TransactionId"`
	UserId        int64  `json:"UserId"`
}

func GetUserInfo(c *fiber.Ctx) UserInfo {
	userInfo := c.Locals(enum.USER_INFO)
	if userInfo != nil {
		return userInfo.(UserInfo)
	}
	return UserInfo{}
}

func InitDefaultUserInfo(ctx *fiber.Ctx) {
	ctx.Locals(enum.USER_INFO, UserInfo{
		TransactionId: uuid.New().String(),
		UserId:        1111})
}

func SetUserInfo(c *fiber.Ctx, userInfo UserInfo) {
	c.Locals(enum.USER_INFO, userInfo)
}
