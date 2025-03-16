package testfiles

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func DoExecuteTx(ctx context.Context, tx *sqlx.Tx, fn func(ctx context.Context, tx *sqlx.Tx) error) error {
	return fn(ctx, tx)
}
