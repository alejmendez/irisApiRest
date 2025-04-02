package migration

import (
	"log"

	model "github.com/alejmendez/goApiRest/app/model"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	log.Println("Migrating database")
	DB.AutoMigrate(&model.User{})
}
