package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfers(t *testing.T) Transfer {
	a1 := createRandomAccount(t)
	a2 := createRandomAccount(t)

	tr, err := testingQueries.CreateTransfers(context.Background(), CreateTransfersParams{
		FromAccountID: a1.ID,
		ToAccountID:   a2.ID,
		Amount:        10,
	})

	require.NoError(t, err)
	require.NotNil(t, tr)
	require.NotEmpty(t, tr)
	require.NotZero(t, tr.ID)

	return tr
}

func TestCreateTransfers(t *testing.T) {
	tr := createRandomTransfers(t)
	deleteAllTransferData(t, tr.ID)
}
func deleteTransferById(t *testing.T, id int64) {

	err := testingQueries.DeleteTansfers(context.Background(), id)
	require.NoError(t, err)

	tr, err := testingQueries.GetTransfers(context.Background(), id)
	require.Error(t, err)
	require.Empty(t, tr)

}
func TestDeleteTansfers(t *testing.T) {
	tr := createRandomTransfers(t)
	deleteAllTransferData(t, tr.ID)
}
func getAllRandomTransfers(t *testing.T) {
	var tranList []Transfer
	for i := 0; i < 3; i++ {
		a := createRandomTransfers(t)
		tranList = append(tranList, a)
	}

	trs, err := testingQueries.GetAllTransfers(context.Background())

	require.NoError(t, err)
	require.NotNil(t, trs)
	require.NotEmpty(t, trs)
	require.GreaterOrEqual(t, len(trs), len(tranList))

	for j := range tranList {
		deleteAllTransferData(t, tranList[j].ID)
	}
}
func TestGetAllTransfers(t *testing.T) {
	getAllRandomTransfers(t)
}
func getRandomTransfers(t *testing.T) {
	tr := createRandomTransfers(t)
	exist, err := testingQueries.GetTransfers(context.Background(), tr.ID)
	require.NoError(t, err)
	require.NotEmpty(t, exist)
	require.NotNil(t, exist)
	require.NotZero(t, exist.ID)
	require.Equal(t, tr, exist)

	deleteAllTransferData(t, exist.ID)
}
func TestGetTransfers(t *testing.T) {
	getRandomTransfers(t)
}
func updateRandomTransfers(t *testing.T) {
	exist := createRandomTransfers(t)
	update, err := testingQueries.UpdateTransfers(context.Background(), UpdateTransfersParams{
		ID:            exist.ID,
		FromAccountID: exist.FromAccountID,
		ToAccountID:   exist.ToAccountID,
		Amount:        22,
	})

	require.NoError(t, err)
	require.NotNil(t, update)
	require.NotEmpty(t, update)
	require.NotZero(t, update.ID)
	require.NotEqual(t, exist, update)
	deleteAllTransferData(t, update.ID)
}
func TestUpdateTransfers(t *testing.T) {
	updateRandomTransfers(t)
}
func deleteAllTransferData(t *testing.T, transferId int64) {
	tr, err := testingQueries.GetTransfers(context.Background(), transferId)
	require.NoError(t, err)
	require.NotEmpty(t, tr)
	require.NotNil(t, tr)
	require.NotZero(t, tr.ID)

	deleteTransferById(t, tr.ID)
	deleteAccountById(t, tr.FromAccountID)
	deleteAccountById(t, tr.ToAccountID)
}
