package sample

import (
	"defaultProjectStructure_sqlc/errors"
	"defaultProjectStructure_sqlc/interfaces"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Controller interfaces.Controller
type sampleController struct {
	sampleService *SampleService
	validator     *validator.Validate
}

func newSampleController(service *SampleService) Controller {
	return sampleController{
		sampleService: service,
		validator:     validator.New(),
	}
}

// GET
// @Summary Gets a sample by ID.
// @Tags Sample
// @Accept json
// @Produce json
// @Param id path int true "Sample ID"
// @Param Authorization header string false "Bearer"
// @Success 200 {object} SampleDto
// @Failure 400 {object} errors.Response
// @Router /samples/{id} [get]
func (sc sampleController) GET(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	sample, err := sc.sampleService.FindById(id)
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	return ctx.Status(fiber.StatusOK).JSON(sample)
}

// POST
// @Summary Create a new sample.
// @Tags Sample
// @Accept json
// @Produce json
// @Param input body SampleDto  true "Sample"
// @Param Authorization header string true "Bearer"
// @Success 201 {object} SampleDto
// @Failure 400 {object} errors.Response
// @Router /samples [post]
func (sc sampleController) POST(ctx *fiber.Ctx) error {
	var sampleDto SampleDto
	if err := ctx.BodyParser(&sampleDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	err := sc.validator.Struct(sampleDto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	// create your model in db
	return ctx.Status(fiber.StatusCreated).JSON(nil)
}

func (sampleController) PUT(ctx *fiber.Ctx) error {
	return nil
}

func (sampleController) PATCH(ctx *fiber.Ctx) error {
	return nil
}

func (sampleController) DELETE(ctx *fiber.Ctx) error {
	return nil
}
