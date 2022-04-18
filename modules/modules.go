package modules

import (
	"github.com/alejmendez/goApiRest/modules/auth"
	"github.com/alejmendez/goApiRest/modules/health"
	"github.com/alejmendez/goApiRest/modules/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializeModules(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	health.Initialize(api)
	auth.Initialize(api)
	users.Initialize(api)
}
