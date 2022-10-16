package model

import (
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        string     `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Username  string     `gorm:"unique_index;not null" json:"username"`
	Email     string     `gorm:"unique_index;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

func (base *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", utils.UUIDv4())
}
