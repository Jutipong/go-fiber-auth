package service

import (
	"auth/app/auth/model"
	"auth/app/auth/repository"
	"auth/pkg/config"
	"auth/pkg/enum"
	"auth/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Login(ctx *fiber.Ctx) (model.Response, error)
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return &service{repo}
}

func (s *service) Login(ctx *fiber.Ctx) (result model.Response, err error) {

	// Validation Model
	var authReq model.Auth
	if err = authReq.Validation(ctx); err != nil {
		return result, err
	}

	// Inquiry in db.
	auth, err := s.repo.Inquiry_Auth(authReq.UserName)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return result, err
	}

	//## Validate password
	if !CheckPasswordHash(authReq.Password, auth) {
		return result, fiber.ErrUnauthorized
	}

	result.Token, err = createToken(&auth.User)
	if err != nil {
		return result, err
	}

	return result, nil
}

func createToken(user *model.User) (string, error) {
	// Get config
	_config := config.Server()
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims[enum.USER_INFO] = utils.JsonSerialize(*user)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(_config.Token_Expire)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(_config.Secret_Key))
	if err != nil {
		return t, fiber.ErrUnauthorized
	}

	return t, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// func CheckPasswordHash(password string, hash string) bool {
func CheckPasswordHash(passReq string, auth model.Auth) bool {
	return bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(passReq+auth.Id)) == nil
}
