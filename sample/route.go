package sample

import "github.com/gofiber/fiber/v2"

func SampleRouter(app fiber.Router, service *SampleService) {
	controller := newSampleController(service)
	app.Get("/samples/:id<int>", controller.GET)
	app.Post("/samples", controller.POST)
	// other sample routes
}
