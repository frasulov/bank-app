package user

import (
	"BankApp/errors"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service   UserService
	validator *validator.Validate
}

func NewUserController(service UserService) UserController {
	validate := validator.New()
	return UserController{
		service:   service,
		validator: validate,
	}
}

// CreateUser
// @Summary Create a new User.
// @Tags User
// @Accept json
// @Produce json
// @Param input body CreateUserInput  true "User"
// @Success 201 {object} CreateUserOutput
// @Failure 400 {object} errors.Response
// @Router /users [post]
func (ac UserController) CreateUser(ctx *fiber.Ctx) error {
	var input CreateUserInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	err := ac.validator.Struct(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	user, err := ac.service.CreateUser(input)
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}
