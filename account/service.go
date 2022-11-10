package account

import (
	db "BankApp/db/sqlc"
	my_errors "BankApp/errors"
	"context"
	"github.com/gofiber/fiber/v2"
)

type AccountService struct {
	repository *db.Repository
}

func GetNewAccountService(repository *db.Repository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (as *AccountService) CreateAccount(input CreateAccountInput) (*AccountOutput, error) {
	account, err := as.repository.CreateAccount(context.Background(), db.CreateAccountParams{
		Owner:    input.Owner,
		Balance:  input.Balance,
		Currency: input.Currency,
	})
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
	}
	return toAccountOutput(account), nil
}

func (as *AccountService) GetAccount(id int64) (*AccountOutput, error) {
	account, err := as.repository.GetAccount(context.Background(), id)
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
	}
	return toAccountOutput(account), nil
}

func (as *AccountService) ListAccounts(param *listAccountParam) ([]*AccountOutput, error) {
	accounts, err := as.repository.GetAccounts(context.Background(), db.GetAccountsParams{
		Limit:  param.PageSize,
		Offset: (param.PageId - 1) * param.PageSize,
	})
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
	}
	return toAccountOutputs(accounts), nil
}
