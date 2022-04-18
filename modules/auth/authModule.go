package auth

import (
	"github.com/alejmendez/goApiRest/modules/auth/router"
	"github.com/gofiber/fiber/v2"
)

func Initialize(api fiber.Router) {
	router.SetupRoutes(api)
}
