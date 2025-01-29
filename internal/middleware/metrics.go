package middleware

import (
	"porsche-api/internal/metrics"

	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Incrémente le compteur pour chaque requête
		metrics.RequestCounter.WithLabelValues(c.Method(), c.Path()).Inc()

		return c.Next()
	}
}
