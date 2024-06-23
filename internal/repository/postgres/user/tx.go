package user

import (
	"context"
	"database/sql"

	"github.com/RoseRocket/xerrs"
	"github.com/jmoiron/sqlx"

	cons "github.com/aaalik/anton-users/internal/constant"
)

func (ur *UserRepository) CreateTx(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := ur.dbw.BeginTxx(ctx, nil)
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLCreateTransaction)
		return nil, err
	}

	return tx, nil
}

func (ur *UserRepository) CommitTx(ctx context.Context, tx *sqlx.Tx) error {
	err := tx.Commit()
	if err != nil {
		err = xerrs.Mask(err, cons.ErrorSQLCommitTransaction)
		return err
	}

	return nil
}

func (ur *UserRepository) RollbackTx(ctx context.Context, tx *sqlx.Tx) error {
	err := tx.Rollback()
	if err != sql.ErrTxDone {
		err = xerrs.Mask(err, cons.ErrorSQLRollbackTransaction)
		return err
	}

	return nil
}
