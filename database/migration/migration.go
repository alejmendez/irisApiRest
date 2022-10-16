package migration

import (
	"log"

	model "github.com/alejmendez/goApiRest/app/models"
	"github.com/jinzhu/gorm"
)

func Migrate(DB *gorm.DB) {
	log.Println("Migrating database")
	DB.AutoMigrate(&model.User{})
}
