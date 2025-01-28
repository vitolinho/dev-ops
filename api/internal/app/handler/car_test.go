package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"path/filepath"
	"porsche-api/internal/domain/model"
	"porsche-api/internal/infrastructure/database"
	"runtime"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Obtenir le chemin de la racine du projet
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../../../")

	// Charger les variables d'environnement de test
	if err := godotenv.Load(filepath.Join(projectRoot, ".env.test")); err != nil {
		log.Fatal("Error loading .env.test file")
	}

	database.Database()
}

func setupTestApp() *fiber.App {
	app := fiber.New()

	// Configuration directe des routes pour les tests
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/cars", AddCar)
	v1.Get("/cars", GetCars)
	v1.Get("/cars/:id", GetCar)
	v1.Put("/cars/:id", UpdateCar)
	v1.Delete("/cars/:id", DeleteCar)

	return app
}

func clearTestData() {
	database.DB.Exec("DELETE FROM cars")
}

func TestAddCar(t *testing.T) {
	clearTestData()
	defer clearTestData()
	app := setupTestApp()

	t.Run("Succès - Ajout d'une voiture", func(t *testing.T) {
		car := model.Car{
			Name:  "Porsche 911",
			Price: 100000,
		}

		body, _ := json.Marshal(car)
		req := httptest.NewRequest("POST", "/api/v1/cars", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	t.Run("Échec - Nom manquant", func(t *testing.T) {
		car := model.Car{
			Price: 100000,
		}

		body, _ := json.Marshal(car)
		req := httptest.NewRequest("POST", "/api/v1/cars", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

func TestGetCars(t *testing.T) {
	clearTestData()
	defer clearTestData()
	app := setupTestApp()

	t.Run("Succès - Récupération de toutes les voitures", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/cars", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

func TestGetCar(t *testing.T) {
	app := setupTestApp()

	t.Run("Succès - Récupération d'une voiture", func(t *testing.T) {
		car := model.Car{
			Name:  "Porsche Cayenne",
			Price: 80000,
		}
		database.DB.Create(&car)

		req := httptest.NewRequest("GET", "/api/v1/cars/"+car.Id.String(), nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Échec - UUID invalide", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/cars/invalid-uuid", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

func TestUpdateCar(t *testing.T) {
	app := setupTestApp()

	t.Run("Succès - Mise à jour d'une voiture", func(t *testing.T) {
		car := model.Car{
			Name:  "Porsche Macan",
			Price: 60000,
		}
		database.DB.Create(&car)

		updateData := model.Car{
			Name:  "Porsche Macan S",
			Price: 70000,
		}

		body, _ := json.Marshal(updateData)
		req := httptest.NewRequest("PUT", "/api/v1/cars/"+car.Id.String(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

func TestDeleteCar(t *testing.T) {
	app := setupTestApp()

	t.Run("Succès - Suppression d'une voiture", func(t *testing.T) {
		car := model.Car{
			Name:  "Porsche Panamera",
			Price: 90000,
		}
		database.DB.Create(&car)

		req := httptest.NewRequest("DELETE", "/api/v1/cars/"+car.Id.String(), nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
