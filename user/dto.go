package user

import "github.com/google/uuid"

type CreateUserInput struct {
	Username string `json:"username" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type LoginUserInput struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}

type LoginUserOutput struct {
	SessionId    uuid.UUID `json:"session_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	User         *UserDto  `json:"user"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token"validate:"required"`
}

type RefreshTokenOutput struct {
	AccessToken string `json:"access_token"`
}
