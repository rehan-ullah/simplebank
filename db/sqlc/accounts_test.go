package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"swiftiesoft.com/simplebank/utils"
)

// create account
func createRandomAccount(t *testing.T) Account {
	args := CreateAccountsParams{
		Owner:    utils.RandomOwner(),
		Balance:  22,
		Currency: utils.RandomCurrency(),
	}

	account, err := testingQueries.CreateAccounts(context.Background(), args)

	require.NoError(t, err)

	require.NotNil(t, account)
	require.NotEmpty(t, account)
	require.NotZero(t, account.ID)

	return account
}
func TestCreateAccounts(t *testing.T) {
	a := createRandomAccount(t)
	deleteAccountById(t, a.ID)
}

// get account
func deleteAccountById(t *testing.T, id int64) {
	err := testingQueries.DeleteAccounts(context.Background(), id)
	require.NoError(t, err)
}
func TestDeleteAccounts(t *testing.T) {
	a := createRandomAccount(t)
	deleteAccountById(t, a.ID)
}

func getAccount(t *testing.T) {
	a := createRandomAccount(t)

	exist, err := testingQueries.GetAccounts(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotNil(t, exist)
	require.NotEmpty(t, exist)
	require.Equal(t, a, exist)

	deleteAccountById(t, a.ID)
}
func TestGetAccounts(t *testing.T) {
	getAccount(t)
}

// get  all accounts

func getAllAccount(t *testing.T) {
	var accoutsList []Account
	limit := 3
	for i := 0; i < limit; i++ {
		a := createRandomAccount(t)
		accoutsList = append(accoutsList, a)
	}

	as, err := testingQueries.GetAllAccounts(context.Background(), GetAllAccountsParams{
		Limit:  int32(limit),
		Offset: 0,
	})

	require.NoError(t, err)
	require.NotNil(t, as)
	require.NotEmpty(t, as)
	require.GreaterOrEqual(t, len(as), len(accoutsList))
	require.LessOrEqual(t, len(as), limit)

	for j := range accoutsList {
		deleteAccountById(t, accoutsList[j].ID)
	}
}

func TestGetAllAccounts(t *testing.T) {
	getAllAccount(t)
}

// update account
func updateAccount(t *testing.T) {
	exist := createRandomAccount(t)

	update, err := testingQueries.UpdateAccounts(context.Background(), UpdateAccountsParams{
		ID:       exist.ID,
		Owner:    utils.RandomOwner(),
		Balance:  44,
		Currency: utils.RandomCurrency(),
	})

	require.NoError(t, err)

	require.NotNil(t, update)
	require.NotEmpty(t, update)
	require.NotZero(t, update.ID)
	require.NotEqual(t, exist, update)
	deleteAccountById(t, update.ID)
}
func TestUpdateAccouts(t *testing.T) {
	updateAccount(t)
}
