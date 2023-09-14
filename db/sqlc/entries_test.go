package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	entry, err := testingQueries.CreateEntries(context.Background(), CreateEntriesParams{
		AccountID: account.ID,
		Amount:    10,
	})

	require.NoError(t, err)

	require.NotNil(t, entry)
	require.NotEmpty(t, entry)
	require.NotZero(t, entry.ID)

	return entry
}
func TestCreateEntries(t *testing.T) {
	entry := createRandomEntry(t)
	deleteAllEntriesData(t, entry.ID)
}

func deleteEntriesById(t *testing.T, id int64) {
	err := testingQueries.DeleteEntries(context.Background(), id)
	require.NoError(t, err)

	entry, err := testingQueries.GetEntries(context.Background(), id)
	require.Error(t, err)
	require.Empty(t, entry)
}
func TestDeleteEntries(t *testing.T) {
	entry := createRandomEntry(t)
	deleteAllEntriesData(t, entry.ID)
}

func getAllRandomEntries(t *testing.T) {

	var entriesList []Entry
	for i := 0; i < 3; i++ {
		a := createRandomEntry(t)
		entriesList = append(entriesList, a)
	}

	ent, err := testingQueries.GetAllEntries(context.Background())

	require.NoError(t, err)
	require.NotNil(t, ent)
	require.NotEmpty(t, ent)
	require.GreaterOrEqual(t, len(ent), len(entriesList))

	for j := range entriesList {
		deleteAllEntriesData(t, entriesList[j].ID)
	}
}
func TestGetAllEntries(t *testing.T) {
	getAllRandomEntries(t)
}
func getRandomEntries(t *testing.T) {
	e := createRandomEntry(t)
	entry, err := testingQueries.GetEntries(context.Background(), e.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotNil(t, entry)
	require.NotZero(t, entry.ID)

	deleteAllEntriesData(t, entry.ID)
}
func TestGetEntries(t *testing.T) {
	getRandomEntries(t)
}

func updateRandomEntries(t *testing.T) {
	entry := createRandomEntry(t)

	update, err := testingQueries.UpdateEntries(context.Background(), UpdateEntriesParams{
		ID:        entry.ID,
		AccountID: entry.AccountID,
		Amount:    20,
	})

	require.NoError(t, err)

	require.NotNil(t, update)
	require.NotEmpty(t, update)
	require.NotZero(t, update.ID)
	require.NotEqual(t, entry, update)

	deleteAllEntriesData(t, update.ID)
}
func TestUpdateEntries(t *testing.T) {
	updateRandomEntries(t)
}

func deleteAllEntriesData(t *testing.T, entryId int64) {
	entry, err := testingQueries.GetEntries(context.Background(), entryId)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotNil(t, entry)
	require.NotZero(t, entry.ID)

	deleteEntriesById(t, entry.ID)
	deleteAccountById(t, entry.AccountID)
}
