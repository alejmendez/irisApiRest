package router

import (
	controller "github.com/alejmendez/goApiRest/modules/auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
}
