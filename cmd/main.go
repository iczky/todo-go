package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/routes"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API is running",
		})
	})

	routes.SetupTodoRoutes(app)

	log.Println("Starting server on :3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
