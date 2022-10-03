package router

import (
	"github.com/alejmendez/goApiRest/app/controllers"
	"github.com/alejmendez/goApiRest/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/health", controllers.Health)

	// Auth
	api.Post("/auth/login", controllers.Login)

	// Users
	api.Get("/users/:id", middleware.Protected(), controllers.GetUser)
	api.Post("/users/", middleware.Protected(), controllers.CreateUser)
	api.Patch("/users/:id", middleware.Protected(), controllers.UpdateUser)
	api.Delete("/users/:id", middleware.Protected(), controllers.DeleteUser)
}
