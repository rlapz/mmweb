package util

import (
	"context"
	"database/sql"

	"github.com/rlapz/mmweb/config"
)

type DbTxExecHandler func(ctx context.Context, trx *sql.Tx, args ...any) error

func DbTxExec(ctx context.Context, db *sql.DB, handler DbTxExecHandler, args ...any) error {
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

func DbTxTryExec(ctx context.Context, conn *sql.DB, query string, args ...any) (int64, error) {
	var rowsCount int64
	xargs := make([]any, 0, 2+len(args))
	xargs = append(xargs, &rowsCount, query)
	xargs = append(xargs, args...)

	err := DbTxExec(ctx, conn, dbTxTryExecHandler, xargs...)
	return rowsCount, err
}

func dbTxTryExecHandler(ctx context.Context, trx *sql.Tx, args ...any) error {
	rc := args[0].(*int64)
	query := args[1].(string)
	xargs := args[2:]

	var err error
	var res sql.Result
	for range config.DB_TRY_MAX {
		res, err = trx.ExecContext(ctx, query, xargs...)
		if err == nil {
			*rc, err = res.RowsAffected()
			break
		}

		if ContextSleep(ctx, config.DB_TRY_WAIT) != nil {
			break
		}
	}

	return err
}
