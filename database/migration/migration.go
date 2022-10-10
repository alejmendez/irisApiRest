package migration

import (
	model "github.com/alejmendez/goApiRest/app/models"
	"github.com/jinzhu/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&model.User{})
}
