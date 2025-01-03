package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/db"
	"github.com/iczky/todo-fiber/internal/handlers"
	"github.com/iczky/todo-fiber/internal/middlewares"
	"github.com/iczky/todo-fiber/internal/models"
	"github.com/iczky/todo-fiber/internal/routes"
	"github.com/iczky/todo-fiber/internal/services"
	"log"
)

func main() {
	app := fiber.New()

	dsn := "postgres://postgres:password@localhost:5432/todo_db"

	database := db.InitDB(dsn)
	defer db.CloseDB()

	log.Println("running database migration...")
	if err := database.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}

	app.Use(middlewares.Logger)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API is running",
		})
	})

	todoService := services.NewTodoService()
	todoHandler := handlers.NewTodoHandler(todoService)

	todoRouter := routes.NewTodoRouter(todoHandler)
	todoRouter.RegisterRoutes(app)

	log.Println("Starting server on :3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
