package utils

import (
	"auth/pkg/enum"

	"github.com/gofiber/fiber/v2"
)

type UserInfo struct {
	TransactionId string `json:"TransactionId"`
	UserId        string `json:"UserId"`
}

func GetUserInfo(c *fiber.Ctx) UserInfo {
	userInfo := c.Locals(enum.USER_INFO)
	if userInfo != nil {
		return userInfo.(UserInfo)
	}
	return UserInfo{}
}

func SetUserInfo(c *fiber.Ctx, userInfo UserInfo) {
	c.Locals(enum.USER_INFO, userInfo)
}
