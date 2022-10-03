package model

import (
	"github.com/alejmendez/goApiRest/core/database"
)

type User struct {
	database.ModelBase
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
