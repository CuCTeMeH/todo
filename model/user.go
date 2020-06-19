package model

import (
	"time"
)

type User struct {
	ID        int        `gorm:"primary_key;column:id"`
	UUID      string     `gorm:"column:uuid"`
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	Email     string     `gorm:"column:email"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
