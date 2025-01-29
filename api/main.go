package main

import (
	"log"
	"net/http"
	"porsche-api/internal/app/router"
	"porsche-api/internal/infrastructure/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Métriques Prometheus
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	database.Database()

	app := fiber.New()
	app.Use(cors.New())

	router.MakeRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
