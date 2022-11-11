package account

import (
	db "BankApp/db/sqlc"
	my_errors "BankApp/errors"
	"context"
	"github.com/gofiber/fiber/v2"
)

type AccountService interface {
	CreateAccount(input CreateAccountInput) (*AccountOutput, error)
	GetAccount(id int64) (*AccountOutput, error)
	ListAccounts(param *ListAccountParam) ([]*AccountOutput, error)
	TransferMoney(param TransferInput) (*db.TransferTxResult, error)
}

type AccountServiceImpl struct {
	repository db.Repository
}

func GetNewAccountService(repository db.Repository) *AccountServiceImpl {
	return &AccountServiceImpl{
		repository: repository,
	}
}

func (as *AccountServiceImpl) CreateAccount(input CreateAccountInput) (*AccountOutput, error) {
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

func (as *AccountServiceImpl) GetAccount(id int64) (*AccountOutput, error) {
	account, err := as.repository.GetAccount(context.Background(), id)
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
	}
	return toAccountOutput(account), nil
}

func (as *AccountServiceImpl) ListAccounts(param *ListAccountParam) ([]*AccountOutput, error) {
	accounts, err := as.repository.GetAccounts(context.Background(), db.GetAccountsParams{
		Limit:  param.PageSize,
		Offset: (param.PageId - 1) * param.PageSize,
	})
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
	}
	return toAccountOutputs(accounts), nil
}

func (as *AccountServiceImpl) TransferMoney(param TransferInput) (*db.TransferTxResult, error) {
	toAccount, err := as.GetAccount(param.ToAccountId)
	if err != nil {
		return nil, err
	}
	if toAccount.Currency != param.Currency {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("currency_not_match", "en"))
	}
	fromAccount, err := as.GetAccount(param.FromAccountId)
	if err != nil {
		return nil, err
	}
	if fromAccount.Currency != param.Currency {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("currency_not_match", "en"))
	}

	result, err := as.repository.TransferTx(context.Background(), db.TransferTxParams{
		FromAccountId: param.FromAccountId,
		ToAccountId:   param.ToAccountId,
		Amount:        param.Amount,
	})
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
	}
	return &result, nil
}
