package test

import (
	"auth/app/auth/model"
	"auth/app/auth/repository"
	"auth/app/auth/service"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_Login(t *testing.T) {
	//Mock Data Customer
	authMock := model.Auth{
		Id:       "1111111111",
		UserName: "root",
		Password: "$2a$12$CzFJXsMQ7Mo0KBAMInbhX.KMVucXTW2Zt/nQrJfPBs65d48VG03nW",
		User:     model.User{UserId: "2222222222"},
	}
	repositoryMock := repository.NewRepositoryMock(&authMock)
	service := service.NewService(repositoryMock)

	req := model.Auth{
		UserName: "root",
		Password: "1234567890",
	}

	result, err := service.Login(&fiber.Ctx{}, &req)

	validateToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVU0VSX0lORk8iOiJ7XCJUcmFuc2FjdGlvbklkXCI6XCI2YjRiNTFhNC04ZGFlLTQ3ZWItYjkzZi05NmJlMmI1NDk2ZGVcIixcIlVzZXJJZFwiOlwiMjIyMjIyMjIyMlwifSIsImV4cCI6MTYzMzQwMzM5MH0.Xw2lTPPSTlm1j97WE9gXkx2MygUjgYp5cqMR0cMkzhs"
	assert.Nil(t, err, nil)
	assert.NotEmpty(t, validateToken, result.Token)
}

func Test_Create(t *testing.T) {
	//Mock Data Customer
	authMock := model.Auth{
		UserName: "root",
		Password: "admin",
	}
	repositoryMock := repository.NewRepositoryMock(&authMock)
	service := service.NewService(repositoryMock)

	req := model.Auth{
		UserName: "root",
		Password: "admin",
	}
	err := service.Create(&fiber.Ctx{}, &req)
	assert.Nil(t, err, nil)
}
