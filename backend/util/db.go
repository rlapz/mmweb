package util

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/rlapz/mmweb/config"
)

type DbTxExecHandler func(ctx context.Context, tx *sql.Tx, args ...any) error

func DbTryExec(ctx context.Context, conn *sql.DB, query string, args ...any) (sql.Result, error) {
	var err error
	var res sql.Result
	for range config.DB_TRY_MAX {
		res, err = conn.ExecContext(ctx, query, args...)
		if err == nil {
			// OK
			break
		}

		if ContextSleep(ctx, config.DB_TRY_WAIT) != nil {
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
		return err
	})

	return ret, err
}

func DbTxTryExecPartial(ctx context.Context, tx *sql.Tx, query string, args ...any) (sql.Result, error) {
	var err error
	var res sql.Result
	for range config.DB_TRY_MAX {
		res, err = tx.ExecContext(ctx, query, args...)
		if err == nil {
			// OK
			break
		}

		if ContextSleep(ctx, config.DB_TRY_WAIT) != nil {
			break
		}
	}

	return res, err
}

func DbTxTryExec(ctx context.Context, conn *sql.DB, query string, args ...any) (sql.Result, error) {
	var ret sql.Result
	var err error
	err = DbTxExecWithHandler(ctx, conn, func(ctx context.Context, tx *sql.Tx, _ ...any) error {
		ret, err = DbTxTryExecPartial(ctx, tx, query, args...)
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

		if ContextSleep(ctx, config.DB_TRY_WAIT) != nil {
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

		if ContextSleep(ctx, config.DB_TRY_WAIT) != nil {
			break
		}
	}

	return ret, err
}

func DbSqlPlaceholder(items any) (string, error) {
	typ := UnwrapPointer(items)
	if typ.Kind() == reflect.Struct {
		return dbSqlPlaceholderSingle(typ)
	}

	if typ.Kind() != reflect.Slice {
		return "", ErrNotSlice
	}

	count := typ.Len()
	if count == 0 {
		return "", ErrSliceZero
	}

	item := typ.Index(0)
	if item.Kind() != reflect.Struct {
		return "", ErrNotStruct
	}

	fcount := item.NumField()
	if fcount == 0 {
		return "", ErrStructZero
	}

	var stb strings.Builder
	stb.Grow((((count * 2) + 1) * fcount) + (count - 1))
	for range count {
		plc := dbSqlPlaceholderBuilder(fcount)
		fmt.Fprintf(&stb, "%s,", plc)
	}

	return stb.String()[:stb.Len()-1], nil
}

func DbDataIsExists(row *sql.Row) (bool, error) {
	var isExists bool
	err := row.Scan(&isExists)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return isExists, nil
}

// Private
func dbSqlPlaceholderBuilder(count int) string {
	var stb strings.Builder
	stb.Grow(1 + (2 * count))
	stb.WriteString("(")
	for range count {
		stb.WriteString("?,")
	}

	return stb.String()[:stb.Len()-1] + ")"
}

func dbSqlPlaceholderSingle(typ reflect.Value) (string, error) {
	count := typ.NumField()
	if count == 0 {
		return "", ErrStructZero
	}

	return dbSqlPlaceholderBuilder(count), nil
}
