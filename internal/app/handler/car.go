package handler

import (
	"fmt"
	"porsche-api/internal/domain/entity"
	"porsche-api/internal/domain/model"
	"porsche-api/internal/infrastructure/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddCar(c* fiber.Ctx) error {
	input := new(model.Car)
	if err := c.BodyParser(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if input.Name == "" || input.Price == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	car := &model.Car{
		Name: input.Name,
		Price: input.Price,
	}

	if result := database.DB.Create(&car); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map { "message" : "Car added.", "data": nil })
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
	return c.JSON(fiber.Map{ "message": "Cars found.", "data": response })
}

func GetCar(c *fiber.Ctx) error {
	id := c.Params("id")
	carID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Id must be an integer.", "data": nil})
	}

	var count int64
	database.DB.Model(&model.Car{}).Count(&count)

	if carID <= 0 || carID > int(count) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Sprintf("Id must be between 1 and %d.", count), "data": nil})
	}

	var car model.Car
	if result := database.DB.Find(&car, id); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{ "message": "Car found.", "data": fiber.Map{ "id": car.Id, "name": car.Name, "price": car.Price } })
}

func UpdateCar(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Id must be an integer.", "data": nil})
	}

	input := new(model.Car)
	if err := c.BodyParser(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var car model.Car
	if result := database.DB.Find(&car, id); result.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Car not found", "data": result.Error})
	}

	if input.Name != "" && input.Name != car.Name {
		car.Name = input.Name
	}

	if input.Price != 0 && input.Price != car.Price {
		car.Price = input.Price
	}

	database.DB.Updates(&car)

	return c.JSON(fiber.Map{ "message": "Car updated.", "data": nil })
}

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Id must be an integer.", "data": nil})
	}

	var car model.Car
	if result := database.DB.Delete(&car, id); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{ "message": "Car deleted.", "data": nil })
}
