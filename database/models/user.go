package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int       `json:"id" gorm:"column:id" gorm:"default:uuid_generate_v3()"`
	Name      string    `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
