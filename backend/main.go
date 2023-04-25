package main

import (
	"github.com/Abhishek-Dobliyal/backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/test", handlers.TestEndpoint)
	app.Post("/", handlers.RetrieveFromApi)
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":5000")
}
