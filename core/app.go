package core

import (
	"fmt"
	"log"
	"sync"

	"github.com/alejmendez/goApiRest/core/config"
	"github.com/alejmendez/goApiRest/core/database"
	"github.com/alejmendez/goApiRest/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	s.app = fiber.New()
	s.InitMiddlewares()

	s.ConnectDB()

	modules.InitializeModules(s.app)

	s.Listen()
}

func (s *Server) ConnectDB() {
	s.DB = database.ConnectDB()
}

func (s *Server) InitMiddlewares() {
	s.app.Use(cors.New())
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
