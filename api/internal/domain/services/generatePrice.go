package services

import "math/rand"

func GeneratePrice() int {
	price := rand.Intn(500000 - 10000) + 10000
	return price
}
