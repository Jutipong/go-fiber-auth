package main

import (
	"auth/pkg/config"
	"auth/pkg/enum"
	"auth/pkg/utils"
	"auth/routes"

	"auth/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.InitialConfig()
	config.InitialDB()
}

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*"}))

	// Middleware
	app.Use(middleware.Logger)

	// Routes
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	_configEnv := config.Server()
	if _configEnv.Env_Mode == enum.ModeDebug {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
