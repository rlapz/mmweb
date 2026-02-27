package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/db"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

type Todo struct {
	db *db.SqlitePool
}

func TodoNew(db *db.SqlitePool) *Todo {
	return &Todo{
		db: db,
	}
}

func (t *Todo) Insert(ctx context.Context, todo *model.Todo, userId int32) error {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	_, err := util.DbTryExec(ctx, conn.Db, query.TodoInsert, userId, todo.Label, todo.CreatedAt)
	return err
}

func (t *Todo) InsertItems(ctx context.Context, items []model.TodoItem) error {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	slcs, err := util.StructToAnySlice(items)
	if err != nil {
		return err
	}

	plc, _ := util.DbSqlPlaceholder(items)
	_, err = util.DbTryExec(ctx, conn.Db, query.TodoInsert+plc, slcs...)
	return err
}

func (t *Todo) IsExists(ctx context.Context, label string, userId int32) (bool, error) {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	row := conn.Db.QueryRowContext(ctx, query.TodoIsExists, label, userId)
	return util.DbDataIsExists(row)
}

func (t *Todo) SelectById(ctx context.Context, id int32) (*model.Todo, error) {
	return nil, nil
}

func (t *Todo) SelectByUserId(ctx context.Context, id int32) ([]model.Todo, error) {
	return nil, nil
}

func (t *Todo) SelectItemsById(ctx context.Context, id int32) ([]model.TodoItem, error) {
	return nil, nil
}
