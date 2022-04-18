package health

import (
	"github.com/alejmendez/goApiRest/modules/health/router"
	"github.com/gofiber/fiber/v2"
)

func Initialize(api fiber.Router) {
	router.SetupRoutes(api)
}
