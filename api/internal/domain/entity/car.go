package entity

import "github.com/google/uuid"

type Car struct {
	Id        uuid.UUID   `json:"id"`
	Name        string `json:"name"`
	Price int `json:"price"`
}
