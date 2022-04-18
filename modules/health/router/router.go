package router

import (
	"github.com/alejmendez/goApiRest/modules/health/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	api.Get("/version", controllers.Version)
}
