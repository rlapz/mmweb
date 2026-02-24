package util

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/errorx"
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

// 'args' must be a struct
// ret:
//
//	0: query placeholder values
//	1: slice of 'args' fields
//	2: error
func SqlPrepareStruct(args any) (string, []any, error) {
	typ := reflect.ValueOf(args)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return "", nil, errorx.DataInvalid
	}

	count := typ.NumField()
	if count == 0 {
		return "", nil, errorx.DataInvalid
	}

	var stb strings.Builder
	stb.Grow(1 + (2 * count))
	stb.WriteString("(")

	nargs := make([]any, 0, count)
	for i := range count {
		val := typ.Field(i).Interface()
		nargs = append(nargs, val)

		stb.WriteString("?,")
	}

	plc := stb.String()
	plc = plc[:stb.Len()-1] + ")"

	return plc, nargs, nil
}

// 'args' must be a slice of structs
// ret:
//
//	0: query placeholder values
//	1: slice of 'args' fields
//	2: error
func SqlPrepareStructSlice(args any) (string, []any, error) {
	typ := reflect.ValueOf(args)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Slice {
		return "", nil, errorx.DataInvalid
	}

	count := typ.Len()

	var stb strings.Builder
	var slcs []any
	for i := range count {
		item := typ.Index(i).Interface()
		plc, slc, err := SqlPrepareStruct(item)
		if err != nil {
			return "", nil, err
		}

		slcs = append(slcs, slc...)
		fmt.Fprintf(&stb, "%s,", plc)
	}

	plcs := stb.String()
	plcs = plcs[:stb.Len()-1]

	return plcs, slcs, nil
}
