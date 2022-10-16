package router

import (
	"github.com/alejmendez/goApiRest/app/controllers"
	"github.com/alejmendez/goApiRest/app/repositories"
	"github.com/alejmendez/goApiRest/app/services"
	"github.com/alejmendez/goApiRest/core/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(app *fiber.App, conf *config.Config, db *gorm.DB) {
	api := app.Group("/api", logger.New())
	healthController := controllers.NewHealthController(conf)
	api.Get("/health", healthController.Get)

	// Repositories
	userRepository := repositories.NewUserRepository(db)

	// Services
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)

	// Controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	// Auth
	api.Post("/auth/login", authController.Login)

	// Users

	userApi := api.Group("/users") //.Use(middleware.Auth)
	userApi.Get("/:id", userController.Get)
	userApi.Post("/", userController.Create)
	userApi.Patch("/:id", userController.Update)
	userApi.Delete("/:id", userController.Delete)

	app.Use(func(ctx *fiber.Ctx) error {
		// Return HTTP 404 status and JSON response.
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "sorry, endpoint is not found",
		})
	},
	)
}
