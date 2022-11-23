package account

import (
	"BankApp/config"
	"BankApp/errors"
	"BankApp/token"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service   AccountService
	validator *validator.Validate
}

func NewAccountController(service AccountService) Controller {
	validate := validator.New()
	_ = validate.RegisterValidation("currency", validateCurrency)
	return Controller{
		service:   service,
		validator: validate,
	}
}

// GetAccount
// @Summary Get an account by ID.
// @Tags Account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param Authorization header string false "Bearer"
// @Success 200 {object} AccountOutput
// @Failure 400 {object} errors.Response
// @Router /accounts/{id} [get]
func (ac Controller) GetAccount(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	account, err := ac.service.GetAccount(int64(id))
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	authPayload := ctx.Locals(config.Configuration.Token.AuthorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		return ctx.Status(fiber.StatusForbidden).JSON(errors.NewResponseByKey("not_permitted", "en"))
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
func (ac Controller) ListAccounts(ctx *fiber.Ctx) error {
	p := new(ListAccountParam)
	if err := ctx.QueryParser(p); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	p.setDefaults()

	authPayload := ctx.Locals(config.Configuration.Token.AuthorizationPayloadKey).(*token.Payload)
	accounts, err := ac.service.ListAccounts(authPayload.Username, p)
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
func (ac Controller) CreateAccount(ctx *fiber.Ctx) error {
	var input CreateAccountInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	err := ac.validator.Struct(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}

	authPayload := ctx.Locals(config.Configuration.Token.AuthorizationPayloadKey).(*token.Payload)
	input.Owner = authPayload.Username
	account, err := ac.service.CreateAccount(input)
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	return ctx.Status(fiber.StatusCreated).JSON(account)
}

// TransferMoney
// @Summary Transfer money between 2 accounts
// @Tags Account
// @Accept json
// @Produce json
// @Param input body TransferInput true "transfer"
// @Success 200 {object} db.TransferTxResult
// @Failure 400 {object} errors.Response
// @Router /transfer [post]
func (ac Controller) TransferMoney(ctx *fiber.Ctx) error {
	var input TransferInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	err := ac.validator.Struct(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewResponseByKey("data_not_valid", "en"))
	}
	result, err := ac.service.TransferMoney(input)
	if err != nil {
		return ctx.Status(err.(errors.HttpError).Code).JSON(err.(errors.HttpError).Response)
	}
	return ctx.Status(fiber.StatusCreated).JSON(result)
}
