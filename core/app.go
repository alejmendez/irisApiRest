package core

import (
	"fmt"
	"log"
	"sync"

	"github.com/alejmendez/goApiRest/core/config"
	"github.com/alejmendez/goApiRest/core/database"
	"github.com/alejmendez/goApiRest/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var once sync.Once

type Server struct {
	app *fiber.App
	DB  *gorm.DB
	api fiber.Router
}

var serverInstance *Server

func GetServerInstance() *Server {
	if serverInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating server instance.")
				serverInstance = &Server{}
			})
	}

	return serverInstance
}

func GetDB() *gorm.DB {
	return GetServerInstance().GetDB()
}

func (s *Server) Start() {
	s.app = fiber.New(fiber.Config{
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
	})

	s.InitMiddlewares()
	s.InitRouter()

	s.ConnectDB()
	s.Listen()
}

func (s *Server) ConnectDB() {
	s.DB = database.ConnectDB()
}

func (s *Server) InitMiddlewares() {
	s.app.Use(cors.New())
	s.app.Use(compress.New())
	s.app.Use(recover.New())
}

func (s *Server) InitRouter() {
	router.SetupRoutes(s.app)
}

func (s *Server) Listen() {
	port := config.Get("APP_PORT")
	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", port)))
}

func (s *Server) GetApp() *fiber.App {
	return s.app
}

func (s *Server) GetDB() *gorm.DB {
	return s.DB
}

func (s *Server) GetRouteApi() fiber.Router {
	return s.api
}

func (s *Server) Close() {
	s.DB.Close()
}
