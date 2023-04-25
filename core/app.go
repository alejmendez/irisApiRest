package core

import (
	"log"

	"github.com/alejmendez/goApiRest/app/utils"
	"github.com/alejmendez/goApiRest/core/config"
	"github.com/alejmendez/goApiRest/core/database"
	"github.com/alejmendez/goApiRest/database/migration"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	App  *fiber.App
	DB   *gorm.DB
	Conf *config.Config
}

func NewServer() (*Server, error) {
	app := &Server{}
	app.Start()

	return app, nil
}

func (s *Server) Start() {
	s.App = fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
}

func (s *Server) Use(args ...interface{}) fiber.Router {
	return s.App.Use(args...)
}

func (s *Server) InitConfig(fileConfEnv string) {
	err := godotenv.Load(fileConfEnv)
	if err != nil {
		log.Print("Error loading .env file")
	}
	s.Conf, _ = config.InitConfig()
}

func (s *Server) ConnectDB() {
	s.DB = database.ConnectDB()
	migration.Migrate(s.DB)
}

func (s *Server) Listen() {
	port := s.Conf.AppPort
	log.Fatal(s.App.Listen(":" + port))
}

func (s *Server) Close() {
	s.DB.Close()
}
