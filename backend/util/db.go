package util

import (
	"context"
	"database/sql"

	"github.com/rlapz/mmweb/config"
)

type DbTxExecHandler func(ctx context.Context, trx *sql.Tx, args ...any) error

func DbTxExec(ctx context.Context, db *sql.DB, query string, args ...any) (int64, error) {
	var ret int64
	err := DbTxExecWithHandler(ctx, db, func(ctx context.Context, trx *sql.Tx, _ ...any) error {
		res, err := trx.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}

		ret, err = res.RowsAffected()
		return err
	})

	return ret, err
}

func DbTxTryExec(ctx context.Context, conn *sql.DB, query string, args ...any) (int64, error) {
	var ret int64
	err := DbTxExecWithHandler(ctx, conn, func(ctx context.Context, trx *sql.Tx, _ ...any) error {
		var err error
		var res sql.Result
		for range config.DB_TRY_MAX {
			res, err = trx.ExecContext(ctx, query, args...)
			if err == nil {
				ret, err = res.RowsAffected()
				break
			}

			if ContextSleep(ctx, config.DB_TRY_WAIT) != nil {
				break
			}
		}

		return err
	})

	return ret, err
}

func DbTxExecWithHandler(ctx context.Context, db *sql.DB, handler DbTxExecHandler, args ...any) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	err = handler(ctx, tx, args...)
	if err != nil {
		return err
	}

	// don't forget to update err variable
	err = tx.Commit()
	return err
}
