package account

import "github.com/gofiber/fiber/v2"

func Router(app fiber.Router, service *AccountService) {
	controller := newAccountController(service)
	app.Get("/accounts/:id<int;min(0)>", controller.GetAccount)
	app.Get("/accounts", controller.ListAccounts)
	app.Post("/accounts", controller.CreateAccount)
	// other sample routes
}
