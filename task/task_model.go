package task

import "time"

type Task struct {
	ID          int        `gorm:"primary_key;column:id"`
	UUID        string     `gorm:"column:uuid"`
	UserID      int        `gorm:"column:user_id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"type:longtext;column:description"`
	Status      string     `gorm:"column:status"`
	Deadline    time.Time  `gorm:"column:deadline"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}
