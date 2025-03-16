package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (db *DB) ExecuteTx(ctx context.Context, tx *sqlx.Tx, fn func(ctx context.Context, tx *sqlx.Tx) error) error {
	if tx != nil {
		return fn(ctx, tx)
	}

	tx, err := db.Writer.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}

	err = fn(ctx, tx)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
