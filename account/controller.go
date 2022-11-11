package account

import (
	"BankApp/errors"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	service   AccountService
	validator *validator.Validate
}

func NewAccountController(service AccountService) AccountController {
	return AccountController{
		service:   service,
		validator: validator.New(),
	}
}

// GetAccount
// @Summary Get a account by ID.
// @Tags Account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param Authorization header string false "Bearer"
// @Success 200 {object} AccountOutput
// @Failure 400 {object} errors.Response
// @Router /accounts/{id} [get]
func (ac AccountController) GetAccount(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	account, err := ac.service.GetAccount(int64(id))
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	return ctx.Status(fiber.StatusOK).JSON(account)
}

// ListAccounts
// @Summary Get accounts
// @Tags Account
// @Accept json
// @Produce json
// @Param page_size query int false "Page Size"
// @Param page_id query int false "Page ID"
// @Param Authorization header string false "Bearer"
// @Success 200 {array} AccountOutput
// @Failure 400 {object} errors.Response
// @Router /accounts [get]
func (ac AccountController) ListAccounts(ctx *fiber.Ctx) error {
	p := new(ListAccountParam)
	if err := ctx.QueryParser(p); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}

	p.setDefaults()
	accounts, err := ac.service.ListAccounts(p)
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	return ctx.Status(fiber.StatusOK).JSON(accounts)
}

// CreateAccount
// @Summary Create a new account.
// @Tags Account
// @Accept json
// @Produce json
// @Param input body CreateAccountInput  true "account"
// @Success 201 {object} AccountOutput
// @Failure 400 {object} errors.Response
// @Router /accounts [post]
func (ac AccountController) CreateAccount(ctx *fiber.Ctx) error {
	var input CreateAccountInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	err := ac.validator.Struct(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	account, err := ac.service.CreateAccount(input)
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	return ctx.Status(fiber.StatusCreated).JSON(account)
}
