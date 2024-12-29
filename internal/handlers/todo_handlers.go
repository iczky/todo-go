package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/services"
	"strconv"
)

type TodoHandler struct {
	Service services.TodoService
}

// Constructor
func NewTodoHandler(service services.TodoService) *TodoHandler {
	return &TodoHandler{
		Service: service,
	}
}

// Generic method to parse request body
func parseRequestBody[T any](c *fiber.Ctx, req *T) error {
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	return nil
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	todos := h.Service.GetAllTodos()
	return c.JSON(todos)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	type request struct {
		Title string `json:"title"`
	}

	body := new(request)
	if err := parseRequestBody(c, body); err != nil {
		return err
	}

	todo := h.Service.CreateTodo(body.Title)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid ID",
			},
		)
	}

	type request struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	body := new(request)
	if err := parseRequestBody(c, body); err != nil {
		return err
	}

	todo, updateErr := h.Service.UpdateTodoByID(id, body.Title, body.Completed)
	if updateErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": updateErr.Error(),
		})
	}

	return c.JSON(todo)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid ID",
			},
		)
	}

	if err := h.Service.DeleteTodoByID(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
