package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

func (r *Repo) TodoInsert(ctx context.Context, todo *model.Todo, userId int32) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	_, err := conn.Db.ExecContext(ctx, query.TodoInsert, userId, todo.Label, todo.CreatedAt)
	return err
}

func (r *Repo) TodoInsertItems(ctx context.Context, items []model.TodoItem) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	slcs, err := util.StructToAnySlice(items)
	if err != nil {
		return err
	}

	plc, _ := util.DbSqlPlaceholder(items)
	res, err := conn.Db.ExecContext(ctx, query.TodoInsert+plc, slcs...)
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

	return err
}

func (r *Repo) TodoIsExists(ctx context.Context, label string, userId int32) (bool, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	row := conn.Db.QueryRowContext(ctx, query.TodoIsExists, label, userId)
	return util.DbDataIsExists(row)
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
