package test

import (
	db "BankApp/db/sqlc"
	"BankApp/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) db.Account {
	user := createRandomUser(t)
	args := db.CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.Equal(t, account2.ID, account1.ID)
	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, time.Second)
}

func TestGetAccounts(t *testing.T) {
	account1 := createRandomAccount(t)
	accounts, err := testQueries.GetAccounts(context.Background(), db.GetAccountsParams{
		Owner:  account1.Owner,
		Limit:  1,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEqual(t, err, sql.ErrNoRows)
	require.NotEqual(t, err, sql.ErrConnDone)
	require.NotEmpty(t, accounts)
	require.Equal(t, 1, len(accounts))
	require.Equal(t, account1.ID, accounts[0].ID)
	require.Equal(t, account1.Currency, accounts[0].Currency)
	require.Equal(t, account1.Owner, accounts[0].Owner)
	require.Equal(t, account1.Balance, accounts[0].Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	delErr := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, delErr)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, account2)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.UpdateAccount(context.Background(), db.UpdateAccountParams{
		account1.ID,
		999,
	})
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, int64(999), account2.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.Equal(t, account2.ID, account1.ID)
}
