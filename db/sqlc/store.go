package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb: err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

/*
 * Delete below code. It is just example
 */

func (store *Store) SampleTransferTx(ctx context.Context, args any) (any, error) {
	err := store.execTx(ctx, func(q *Queries) error {
		// call your repository functions here
		// it is example generally we use insert/update in transactions
		q.GetSampleById(ctx, 1)
		q.GetSamples(ctx)
		return nil
	})
	return nil, err
}
