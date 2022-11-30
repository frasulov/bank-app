package user

import (
	db "BankApp/db/sqlc"
	"github.com/google/uuid"
)

func toUserOutput(model db.User) *UserDto {
	return &UserDto{
		Username: model.Username,
		FullName: model.FullName,
		Email:    model.Email,
	}
}

func toLoginResponseOutput(accessToken, refreshToken string, sessionId uuid.UUID, user db.User) *LoginUserOutput {
	return &LoginUserOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		SessionId:    sessionId,
		User:         toUserOutput(user),
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
