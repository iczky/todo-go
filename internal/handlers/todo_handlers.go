package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/todo-fiber/internal/services"
	"strconv"
)

func GetTodos(c *fiber.Ctx) error {
	todos := services.GetAllTodos()
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	type request struct {
		Title string `json:"title"`
	}

	body := new(request)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	todo := services.CreateTodo(body.Title)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
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
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	todo, updateErr := services.UpdateTodoByID(id, body.Title, body.Completed)
	if updateErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": updateErr.Error(),
		})
	}

	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid ID",
			},
		)
	}

	if err := services.DeleteTodoByID(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
