package router

import (
	controller "github.com/alejmendez/goApiRest/handler"
	"github.com/alejmendez/goApiRest/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", controller.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", controller.GetUser)
	user.Post("/", controller.CreateUser)
	user.Patch("/:id", middleware.Protected(), controller.UpdateUser)
	user.Delete("/:id", middleware.Protected(), controller.DeleteUser)

	// Product
	product := api.Group("/product")
	product.Get("/", controller.GetAllProducts)
	product.Get("/:id", controller.GetProduct)
	product.Post("/", middleware.Protected(), controller.CreateProduct)
	product.Delete("/:id", middleware.Protected(), controller.DeleteProduct)
}
