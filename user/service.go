package user

import (
	"BankApp/config"
	db "BankApp/db/sqlc"
	my_errors "BankApp/errors"
	"BankApp/globals"
	"BankApp/util"
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type UserService interface {
	CreateUser(input CreateUserInput) (*UserDto, error)
	LoginUser(data LoginUserInput) (*LoginUserOutput, error)
}

type UserServiceImpl struct {
	repository db.Repository
}

func GetNewUserService(repository db.Repository) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (us *UserServiceImpl) CreateUser(input CreateUserInput) (*UserDto, error) {
	var err error
	input.Password, err = util.HashPassword(input.Password)
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	user, err := us.repository.CreateUser(context.Background(), toUserModel(input))
	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code.Name() {
		case "unique_violation":
			return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("username_or_email_exist", "en"))
		default:
			return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("system_error", "en"))
		}
	}
	return toUserOutput(user), nil
}

func (us *UserServiceImpl) LoginUser(data LoginUserInput) (*LoginUserOutput, error) {
	user, err := us.repository.GetUser(context.Background(), data.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
		}
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	err = util.CheckPassword(user.Password, data.Password)
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusUnauthorized, my_errors.NewResponseByKey("credentials_error", "en"))
	}

	accessToken, err := globals.TokenMaker.CreateToken(data.Username, config.Configuration.AccessTokenDuration)
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	return toLoginResponseOutput(accessToken, user), nil
}
