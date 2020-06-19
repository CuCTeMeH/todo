package model

import (
	"time"
)

type User struct {
	ID        int        `gorm:"primary_key;column:id"`
	UUID      string     `gorm:"column:uuid"`
	Username  string     `gorm:"column:username"`
	Email     string     `gorm:"column:email"`
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
