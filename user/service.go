package user

import (
	"BankApp/config"
	db "BankApp/db/sqlc"
	my_errors "BankApp/errors"
	"BankApp/globals"
	"BankApp/util"
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, input CreateUserInput) (*UserDto, error)
	LoginUser(ctx context.Context, data LoginUserInput) (*LoginUserOutput, error)
	RenewToken(ctx context.Context, data RefreshTokenInput) (*RefreshTokenOutput, error)
}

type UserServiceImpl struct {
	repository db.Repository
}

func GetNewUserService(repository db.Repository) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (us *UserServiceImpl) CreateUser(ctx context.Context, input CreateUserInput) (*UserDto, error) {
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

func (us *UserServiceImpl) RenewToken(ctx context.Context, data RefreshTokenInput) (*RefreshTokenOutput, error) {

	refreshTokenPayload, err := globals.TokenMaker.VerifyToken(data.RefreshToken)
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusUnauthorized, my_errors.NewResponseByKey("refresh_token_not_valid", "en"))
	}

	session, err := us.repository.GetSession(ctx, refreshTokenPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
		}
		log.Printf("ERR while getting user: %s", err.Error())
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	if session.IsBlocked {
		return nil, my_errors.NewHttpError(fiber.StatusUnauthorized, my_errors.NewResponseByKey("blocked_session", "en"))
	}

	if session.Username != refreshTokenPayload.Username {
		return nil, my_errors.NewHttpError(fiber.StatusUnauthorized, my_errors.NewResponseByKey("incorrect_session_user", "en"))
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, my_errors.NewHttpError(fiber.StatusUnauthorized, my_errors.NewResponseByKey("expired_session", "en"))
	}

	accessToken, _, err := globals.TokenMaker.CreateToken(session.Username, config.Configuration.AccessTokenDuration)
	if err != nil {
		log.Printf("ERR while creating user: %s", err.Error())
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	return &RefreshTokenOutput{
		AccessToken: accessToken,
	}, nil

}

func (us *UserServiceImpl) LoginUser(ctx context.Context, data LoginUserInput) (*LoginUserOutput, error) {
	user, err := us.repository.GetUser(ctx, data.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
		}
		log.Printf("ERR while getting user: %s", err.Error())
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	err = util.CheckPassword(user.Password, data.Password)
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusUnauthorized, my_errors.NewResponseByKey("credentials_error", "en"))
	}

	accessToken, _, err := globals.TokenMaker.CreateToken(data.Username, config.Configuration.AccessTokenDuration)
	if err != nil {
		log.Printf("ERR while creating user: %s", err.Error())
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	refrseshToken, payload, err := globals.TokenMaker.CreateToken(data.Username, config.Configuration.RefreshTokenDuration)
	if err != nil {
		log.Printf("ERR while creating user: %s", err.Error())
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	fmt.Println("string(ctx.(*fasthttp.RequestCtx).UserAgent()): ", string(ctx.(*fasthttp.RequestCtx).RemoteIP()))
	userSession, err := us.repository.CreateSession(context.Background(), db.CreateSessionParams{
		ID:           payload.ID,
		Username:     user.Username,
		RefreshToken: refrseshToken,
		UserAgent:    string(ctx.(*fasthttp.RequestCtx).UserAgent()),
		ClientIp:     ctx.(*fasthttp.RequestCtx).RemoteIP().String(),
		IsBlocked:    false,
		ExpiresAt:    payload.ExpiredAt,
	})
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusInternalServerError, my_errors.NewResponseByKey("system_error", "en"))
	}

	return toLoginResponseOutput(accessToken, refrseshToken, userSession.ID, user), nil
}
