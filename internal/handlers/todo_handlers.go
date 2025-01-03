package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/services"
	"gorm.io/gorm"
	"strconv"
)

type TodoHandler struct {
	Service services.TodoService
}

type CreateTodoRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Priority    int     `json:"priority" validate:"required,min=1,max=5"`
	DueDate     *string `json:"due_date" validate:"omitempty,datetime=2006-01-02"`
}

type UpdateTodoRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Priority    int     `json:"priority"`
	DueDate     *string `json:"due_date"`
	Completed   bool    `json:"completed"`
}

// NewTodoHandler Constructor
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
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": todos,
	})
}

func (h *TodoHandler) GetTodoById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch todo",
		})
	}
	todo, err := h.Service.GetTodoById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": todo,
	})
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	body := new(CreateTodoRequest)
	if err := parseRequestBody(c, body); err != nil {
		return err
	}

	todo, err := h.Service.CreateTodo(body.Title, body.Description, body.Priority, body.DueDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

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

	body := new(UpdateTodoRequest)
	if err := parseRequestBody(c, body); err != nil {
		return err
	}

	todo, updateErr := h.Service.UpdateTodoByID(id, body.Title, body.Description, body.Completed, body.Priority, body.DueDate)
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

	deleteErr := h.Service.DeleteTodoByID(id)
	if deleteErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": deleteErr.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
