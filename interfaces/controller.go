package interfaces

import "github.com/gofiber/fiber/v2"

type Controller interface {
	GET(c *fiber.Ctx) error
	POST(c *fiber.Ctx) error
	PUT(c *fiber.Ctx) error
	DELETE(c *fiber.Ctx) error
	PATCH(c *fiber.Ctx) error
}
