package handler

import (
	"porsche-api/internal/domain/entity"
	"porsche-api/internal/domain/model"
	"porsche-api/internal/infrastructure/database"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

func AddCar(c* fiber.Ctx) error {
	input := new(model.Car)
	if err := c.BodyParser(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if input.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON("name is required")
	}

	if input.Price == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("price cannot be 0")
	}

	car := &model.Car{
		Name: input.Name,
		Price: input.Price,
	}

	if result := database.DB.Create(&car); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON("Car added")
}

func GetCars(c *fiber.Ctx) error {
	var cars []model.Car
	if result := database.DB.Find(&cars); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	response := make([]entity.Car, len(cars))
	for i, r := range cars {
		response[i] = entity.Car{
			Id: r.Id,
			Name: r.Name,
			Price: r.Price,
		}
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetCar(c *fiber.Ctx) error {
	id := c.Params("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid UUID format")
	}

	var car model.Car
	if result := database.DB.Where("id = ?", parsedID).First(&car); result.Error != nil {
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON("Car not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to retrieve car")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":    car.Id,
		"name":  car.Name,
		"price": car.Price,
	})
}

func UpdateCar(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid UUID format")
	}

	input := new(model.Car)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid request body")
	}

	var car model.Car
	if result := database.DB.Where("id = ?", parsedID).First(&car); result.Error != nil {
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON("Car not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to fetch car")
	}

	if input.Name != "" && input.Name != car.Name {
		car.Name = input.Name
	}

	if input.Price != 0 && input.Price != car.Price {
		car.Price = input.Price
	}

	if result := database.DB.Save(&car); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to update car")
	}

	return c.Status(fiber.StatusOK).JSON("Car updated")
}

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid UUID format")
	}

	var car model.Car
	if result := database.DB.Where("id = ?", parsedID).First(&car); result.Error != nil {
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON("Car not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to fetch car")
	}

	if result := database.DB.Delete(&car); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to delete car")
	}

	return c.Status(fiber.StatusOK).JSON("Car deleted")
}
