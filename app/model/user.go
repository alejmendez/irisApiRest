package model

import (
	"time"
)

type User struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"unique_index;not null" json:"username"`
	Email     string     `gorm:"unique_index;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
