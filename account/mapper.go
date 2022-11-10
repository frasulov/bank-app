package account

import db "BankApp/db/sqlc"

func toAccountOutput(model db.Account) *AccountOutput {
	return &AccountOutput{
		Id:       model.ID,
		Currency: model.Currency,
		Balance:  model.Balance,
		Owner:    model.Owner,
	}
}

func toAccountOutputs(models []db.Account) (result []*AccountOutput) {

	for _, value := range models {
		result = append(result, toAccountOutput(value))
	}
	return
}
