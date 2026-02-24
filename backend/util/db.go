package util

import (
	"context"
	"database/sql"

	"github.com/rlapz/mmweb/config"
)

type DbTxExecHandler func(ctx context.Context, tx *sql.Tx, args ...any) error

func DbTxTryPartialExec(ctx context.Context, tx *sql.Tx, query string, args ...any) (sql.Result, error) {
	var err error
	var res sql.Result
	for range config.DB_TRY_MAX {
		res, err = tx.ExecContext(ctx, query, args...)
		if err == nil {
			// OK
			break
		}

		err = ContextSleep(ctx, config.DB_TRY_WAIT)
		if err != nil {
			break
		}
	}

	return res, err
}

func DbTxExec(ctx context.Context, db *sql.DB, query string, args ...any) (sql.Result, error) {
	var ret sql.Result
	var err error
	err = DbTxExecWithHandler(ctx, db, func(ctx context.Context, tx *sql.Tx, _ ...any) error {
		ret, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}

		return nil
	})

	return ret, err
}

func DbTxTryExec(ctx context.Context, conn *sql.DB, query string, args ...any) (sql.Result, error) {
	var ret sql.Result
	var err error
	err = DbTxExecWithHandler(ctx, conn, func(ctx context.Context, tx *sql.Tx, _ ...any) error {
		for range config.DB_TRY_MAX {
			ret, err = tx.ExecContext(ctx, query, args...)
			if err == nil {
				return nil
			}

			err = ContextSleep(ctx, config.DB_TRY_WAIT)
			if err != nil {
				return err
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

func DbTryRowsAffected(ctx context.Context, res sql.Result) (int64, error) {
	var ret int64
	var err error
	for range config.DB_TRY_MAX {
		ret, err = res.RowsAffected()
		if err == nil {
			// OK
			break
		}

		err = ContextSleep(ctx, config.DB_TRY_WAIT)
		if err != nil {
			break
		}
	}

	return ret, err
}

func DbTryLastInsertId(ctx context.Context, res sql.Result) (int64, error) {
	var ret int64
	var err error
	for range config.DB_TRY_MAX {
		ret, err = res.LastInsertId()
		if err == nil {
			// OK
			break
		}

		err = ContextSleep(ctx, config.DB_TRY_WAIT)
		if err != nil {
			break
		}
	}

	return ret, err
}
