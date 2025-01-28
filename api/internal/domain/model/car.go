package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Price int `gorm:"not null" json:"price"`
}
