package model

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Id uint `gorm:"not null" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Price int `gorm:"not null" json:"price"`
}
