package model

import "time"

type List struct {
	ID        int        `gorm:"primary_key;column:id"`
	UUID      string     `gorm:"column:uuid"`
	UserID    int        `gorm:"column:user_id"`
	Name      string     `gorm:"column:name"`
	Status    string     `gorm:"column:status"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
