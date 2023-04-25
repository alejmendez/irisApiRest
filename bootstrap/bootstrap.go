package bootstrap

import (
	"github.com/alejmendez/goApiRest/core"
	"github.com/alejmendez/goApiRest/router"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApplication() *core.Server {
	app, _ := core.NewServer()

	app.InitConfig(".env")
	app.ConnectDB()

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(recover.New())
	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen()

	return app
}
