package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ValidateTitle(c *fiber.Ctx) error {
	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	title, exists := body["title"]
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing title",
		})
	}

	if strings.TrimSpace(title.(string)) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title cannot be empty",
		})
	}

	return c.Next()
}
