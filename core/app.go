package core

import (
	"fmt"
	"log"

	"github.com/alejmendez/goApiRest/core/config"
	"github.com/alejmendez/goApiRest/core/database"
	"github.com/alejmendez/goApiRest/database/migration"
	"github.com/alejmendez/goApiRest/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	app  *fiber.App
	DB   *gorm.DB
	conf *config.Config
}

func fiberConfig() fiber.Config {
	return fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			fmt.Println(err)
			err = ctx.Status(code).JSON(fiber.Map{"status": "error", "message": err.Error()})
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal Server Error"})
			}

			// Return from handler
			return nil
		},
	}
}

func NewServer() (*Server, error) {
	app := &Server{}
	app.Start()

	return app, nil
}

func (s *Server) Start() {
	s.app = fiber.New(fiberConfig())

	s.InitMiddlewares()
	s.InitRouter()
	s.InitConfig()

	s.ConnectDB()
	s.Listen()
}

func (s *Server) InitMiddlewares() {
	s.app.Use(cors.New())
	s.app.Use(compress.New())
	s.app.Use(recover.New())
}

func (s *Server) InitRouter() {
	router.SetupRoutes(s.app)
}

func (s *Server) InitConfig() {
	s.conf, _ = config.InitConfig()
}

func (s *Server) ConnectDB() {
	s.DB = database.ConnectDB()
	migration.Migrate(s.DB)
}

func (s *Server) Listen() {
	port := s.conf.AppPort
	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", port)))
}

func (s *Server) Close() {
	s.DB.Close()
}
