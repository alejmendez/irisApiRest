package database

import (
	"time"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/jinzhu/gorm"
)

type ModelBase struct {
	ID        string `gorm:"unique_index;not null" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *ModelBase) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", utils.UUIDv4())
}
