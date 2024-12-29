package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/handlers"
	"github.com/iczky/todo-fiber/internal/middlewares"
	"github.com/iczky/todo-fiber/internal/routes"
	"github.com/iczky/todo-fiber/internal/services"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(middlewares.Logger)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API is running",
		})
	})

	todoService := services.NewTodoService()
	todoHandler := handlers.NewTodoHandler(todoService)

	routes.SetupTodoRoutes(app, todoHandler)

	log.Println("Starting server on :3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
