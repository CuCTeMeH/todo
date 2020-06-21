package model

import (
	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}
func AutoMigrate() {
	Client().AutoMigrate(&User{}, &List{}, &Task{})
}
