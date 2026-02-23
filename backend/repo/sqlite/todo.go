package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

func (r *Repo) TodoInsert(ctx context.Context, todo *model.Todo, userId int32) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	return util.DbTxExec(ctx, conn.Db, func(ctx context.Context, trx *sql.Tx, args ...any) error {
		res, err := trx.ExecContext(ctx, query.TodoInsert, userId, todo.Label, todo.CreatedAt)
		if err != nil {
			return err
		}

		aff, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if aff == 0 {
			return errorx.NoDataSaved
		}
		return nil
	})
}

func (r *Repo) TodoInsertItems(ctx context.Context, id int32, items []model.TodoItem) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	return util.DbTxExec(ctx, conn.Db, func(ctx context.Context, trx *sql.Tx, args ...any) error {
		// TODO: batch insert
		for i := range items {
			u := &items[i]
			res, err := trx.ExecContext(ctx, query.TodoInsert, id, u.Title, u.Description,
				u.Flags, u.CreatedAt)
			if err != nil {
				return err
			}

			aff, err := res.RowsAffected()
			if err != nil {
				return err
			}

			if aff == 0 {
				return errorx.NoDataSaved
			}
		}

		return nil
	})
}

func (r *Repo) TodoIsExists(ctx context.Context, label string, userId int32) (bool, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	var ok bool
	row := conn.Db.QueryRowContext(ctx, query.TodoIsExists, label, userId)
	err := row.Scan(&ok)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, err
}

func (r *Repo) TodoSelectById(ctx context.Context, id int32) (*model.Todo, error) {
	return nil, nil
}

func (r *Repo) TodoSelectByUserId(ctx context.Context, id int32) ([]model.Todo, error) {
	return nil, nil
}

func (r *Repo) TodoSelectItemsById(ctx context.Context, id int32) ([]model.TodoItem, error) {
	return nil, nil
}
