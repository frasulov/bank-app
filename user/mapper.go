package user

import db "BankApp/db/sqlc"

func toUserOutput(model db.User) *UserDto {
	return &UserDto{
		Username: model.Username,
		FullName: model.FullName,
		Email:    model.Email,
	}
}

func toLoginResponseOutput(accessToken string, user db.User) *LoginUserOutput {
	return &LoginUserOutput{
		AccessToken: accessToken,
		User:        toUserOutput(user),
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
