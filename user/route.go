package user

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router, service *UserServiceImpl) {
	controller := NewUserController(service)
	app.Post("/users", controller.CreateUser)
	app.Post("/login", controller.LoginUser)
	app.Post("/refresh-token", controller.RefreshToken)
	// other sample routes
}
