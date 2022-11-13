package user

import db "BankApp/db/sqlc"

func toUserOutput(model db.User) *CreateUserOutput {
	return &CreateUserOutput{
		Username: model.Username,
		FullName: model.FullName,
		Email:    model.Email,
	}
}

func toUserModel(model CreateUserInput) db.CreateUserParams {
	return db.CreateUserParams{
		Username: model.Username,
		FullName: model.FullName,
		Email:    model.Email,
		Password: model.Password,
	}
}
