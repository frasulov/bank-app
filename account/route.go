package account

import (
	"BankApp/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router, service *AccountServiceImpl) {
	controller := NewAccountController(service)
	app.Get("/accounts/:id<int;min(0)>", middleware.Protect, controller.GetAccount)
	app.Get("/accounts", middleware.Protect, controller.ListAccounts)
	app.Post("/accounts", middleware.Protect, controller.CreateAccount)
	app.Post("/transfer", controller.TransferMoney)
	// other sample routes
}
