package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/handlers"
	"github.com/iczky/todo-fiber/internal/middlewares"
)

type TodoRouter struct {
	handler *handlers.TodoHandler
}

func NewTodoRouter(handler *handlers.TodoHandler) *TodoRouter {
	return &TodoRouter{
		handler: handler,
	}
}

func (r *TodoRouter) RegisterRoutes(app *fiber.App) {
	todo := app.Group("/api/todos")

	todo.Get("/", r.handler.GetTodos)
	todo.Get("/:id", r.handler.GetTodoById)
	todo.Post("/", middlewares.ValidateTitle, r.handler.CreateTodo)
	todo.Put("/:id", middlewares.ValidateTitle, r.handler.UpdateTodo)
	todo.Delete("/:id", r.handler.DeleteTodo)
}
