package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/handlers"
	"github.com/iczky/todo-fiber/internal/middlewares"
)

func SetupTodoRoutes(app *fiber.App, handler *handlers.TodoHandler) {
	todo := app.Group("/api/todos")

	todo.Get("/", handler.GetTodos)
	todo.Post("/", middlewares.ValidateTitle, handler.CreateTodo)
	todo.Put("/:id", middlewares.ValidateTitle, handler.UpdateTodo)
	todo.Delete("/:id", handler.DeleteTodo)
}
