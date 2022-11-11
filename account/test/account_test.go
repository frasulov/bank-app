package test

import (
	"BankApp/account"
	mockdb "BankApp/db/mock"
	my_errors "BankApp/errors"
	"BankApp/util"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestGetAccountAPI200(t *testing.T) {
	accountInstance := randomAccountOutput()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockdb.NewMockService(ctrl)
	service.EXPECT().
		GetAccount(gomock.Eq(accountInstance.Id)).
		Times(1).
		Return(accountInstance, nil)
	controller := account.NewAccountController(service)
	app := fiber.New()
	app.Get("/accounts/:id", controller.GetAccount)

	req := httptest.NewRequest("GET", fmt.Sprintf("/accounts/%v", accountInstance.Id), nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	fmt.Println("resp: ", resp)
}

func TestGetAccountAPI404(t *testing.T) {
	accountInstance := randomAccountOutput()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockdb.NewMockService(ctrl)
	service.EXPECT().
		GetAccount(gomock.Eq(accountInstance.Id)).
		Times(1).
		Return(nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en")))
	controller := account.NewAccountController(service)
	app := fiber.New()
	app.Get("/accounts/:id", controller.GetAccount)

	req := httptest.NewRequest("GET", fmt.Sprintf("/accounts/%v", accountInstance.Id), nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestGetAccountAPI404Query(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mockdb.NewMockService(ctrl)
	controller := account.NewAccountController(service)
	app := fiber.New()
	app.Get("/accounts/:id", controller.GetAccount)

	req := httptest.NewRequest("GET", fmt.Sprintf("/accounts/%v", "dsds"), nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func randomAccountOutput() *account.AccountOutput {
	return &account.AccountOutput{
		Id:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
