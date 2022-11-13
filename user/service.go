package user

import (
	db "BankApp/db/sqlc"
	my_errors "BankApp/errors"
	"BankApp/util"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type UserService interface {
	CreateUser(input CreateUserInput) (*CreateUserOutput, error)
}

type UserServiceImpl struct {
	repository db.Repository
}

func GetNewUserService(repository db.Repository) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (us *UserServiceImpl) CreateUser(input CreateUserInput) (*CreateUserOutput, error) {
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
