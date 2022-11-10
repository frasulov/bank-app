package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Common() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Request().Header.Add("Content-Type", "application/json")
		return c.Next()
	}
}

func Protect() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// add your middleware
		return c.Next()
	}
}
