package main

import (
	"porsche-api/internal/infrastructure/database"
	"porsche-api/internal/app/router"
	"github.com/gofiber/fiber/v2"
	"log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Database()

	app := fiber.New()
	app.Use(cors.New())

	router.MakeRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
