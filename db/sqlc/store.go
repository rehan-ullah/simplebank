package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provice all the functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore create new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a func within  a database transaction
func (store *Store) execTx(c context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams that contain the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAcccountID int64 `json:"from_account_id"`
	ToAccountID    int64 `json:"to_account_id"`
	Amount         int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer     Transfer `json:"transfer"`
	FromAcccount Account  `json:"from_account"`
	ToAccount    Account  `json:"to_account"`
	FromEntry    Entry    `json:"from_entry"`
	ToEntry      Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(c context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(c, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfers(c, CreateTransfersParams{
			FromAccountID: arg.FromAcccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntries(c, CreateEntriesParams{
			AccountID: arg.FromAcccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntries(c, CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
