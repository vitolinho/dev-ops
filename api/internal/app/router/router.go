package router

import (
	"porsche-api/internal/app/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MakeRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	v1 := api.Group("/v1")
	v1.Post("/cars", handler.AddCar)
	v1.Get("/cars", handler.GetCars)
	v1.Get("/cars/:id", handler.GetCar)
	v1.Put("/cars/:id", handler.UpdateCar)
	v1.Delete("/cars/:id", handler.DeleteCar)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
