package users

import (
	"fmt"

	"github.com/alejmendez/goApiRest/core/database"
	models "github.com/alejmendez/goApiRest/modules/users/models"
	"github.com/alejmendez/goApiRest/modules/users/router"
	"github.com/gofiber/fiber/v2"
)

func Initialize(api fiber.Router) {
	router.SetupRoutes(api)
	MigrateModels()
}

func MigrateModels() {
	db := database.DB
	db.AutoMigrate(&models.User{})
	fmt.Println("Database Model User Migrated")
}
