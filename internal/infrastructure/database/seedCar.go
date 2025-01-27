package database

import (
	"porsche-api/internal/domain/model"
	"porsche-api/internal/domain/services"
)

func SeedCar() {
	var count int64
	DB.Model(&model.Car{}).Count(&count)

	if count == 0 {
		cars := []model.Car{
			{ Name: "Porsche 911", Price: services.GeneratePrice() },
			{ Name: "Porsche Cayenne", Price: services.GeneratePrice() },
			{ Name: "Porsche Panamera", Price: services.GeneratePrice() },
			{ Name: "Porsche Macan", Price: services.GeneratePrice() },
			{ Name: "Porsche Boxster", Price: services.GeneratePrice() },
			{ Name: "Porsche Cayman", Price: services.GeneratePrice() },
			{ Name: "Porsche Taycan", Price: services.GeneratePrice() },
			{ Name: "Porsche 718", Price: services.GeneratePrice() },
			{ Name: "Porsche Carrera GT", Price: services.GeneratePrice() },
			{ Name: "Porsche 959", Price: services.GeneratePrice() },
			{ Name: "Porsche 356", Price: services.GeneratePrice() },
			{ Name: "Porsche 550 Spyder", Price: services.GeneratePrice() },
			{ Name: "Porsche 944", Price: services.GeneratePrice() },
			{ Name: "Porsche 928", Price: services.GeneratePrice() },
			{ Name: "Porsche 918 Spyder", Price: services.GeneratePrice() },
			{ Name: "Porsche 912", Price: services.GeneratePrice() },
			{ Name: "Porsche 930", Price: services.GeneratePrice() },
			{ Name: "Porsche 964", Price: services.GeneratePrice() },
			{ Name: "Porsche 993", Price: services.GeneratePrice() },
			{ Name: "Porsche 996", Price: services.GeneratePrice() },
			{ Name: "Porsche 997", Price: services.GeneratePrice() },
			{ Name: "Porsche 991", Price: services.GeneratePrice() },
			{ Name: "Porsche 992", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Cayman", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Boxster", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Cayman GT4", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Boxster Spyder", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Cayman GTS", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Boxster GTS", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Cayman GT4 RS", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Boxster Spyder RS", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Cayman GT4 RS Clubsport", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Boxster Spyder RS Clubsport", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 RSK", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 W-RS", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 Formula 2", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 RSK Spyder", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 RS 60 Spyder", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 RS 61 Spyder", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 GTR", Price: services.GeneratePrice() },
			{ Name: "Porsche 718/2", Price: services.GeneratePrice() },
			{ Name: "Porsche 718/8 W-RS", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 F2", Price: services.GeneratePrice() },
			{ Name: "Porsche 718/8 F1", Price: services.GeneratePrice() },
			{ Name: "Porsche 718/2 1500RS", Price: services.GeneratePrice() },
			{ Name: "Porsche 718/2 F2", Price: services.GeneratePrice() },
			{ Name: "Porsche 718/2 F1", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 RS 60", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 RS 61", Price: services.GeneratePrice() },
			{ Name: "Porsche 718 GTR RS", Price: services.GeneratePrice() },
		}
		for _, car := range cars {
			DB.Create(&car)
		}
	}
}
