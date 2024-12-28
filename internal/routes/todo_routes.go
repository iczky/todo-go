package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/handlers"
)

func SetupTodoRoutes(app *fiber.App) {
	todo := app.Group("/api/todos")

	todo.Get("/", handlers.GetTodos)
	todo.Post("/", handlers.CreateTodo)
	todo.Put("/:id", handlers.UpdateTodo)
	todo.Delete("/:id", handlers.DeleteTodo)
}
