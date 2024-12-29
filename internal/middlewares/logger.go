package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	log.Printf("Time: %s | Method: %s | Path: %s | Status: %d | Duration: %s\n",
		start.Format(time.RFC1123),
		c.Method(),
		c.Path(),
		c.Response().StatusCode(),
		time.Since(start))

	return err
}
