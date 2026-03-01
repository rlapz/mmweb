package sqlite

import (
	"context"
	"database/sql"

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

func (t *Todo) SelectById(ctx context.Context, id int32) (*model.Todo, error) {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	ret := new(model.Todo)
	row := conn.Db.QueryRowContext(ctx, query.TodoSelectById, id)
	err := row.Scan(&ret.Id, &ret.IdUser, &ret.Label, &ret.CreatedAt, &ret.UpdatedAt)
	return ret, err
}

func (t *Todo) SelectByUserId(ctx context.Context, id int32) ([]model.Todo, error) {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	rows, err := conn.Db.QueryContext(ctx, query.TodoSelectByUserId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []model.Todo
	for rows.Next() {
		var item model.Todo
		err = rows.Scan(&item.Id, &item.IdUser, &item.Label, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ret = append(ret, item)
	}

	return ret, nil
}

func (t *Todo) SelectItemById(ctx context.Context, id int32) (*model.TodoItem, error) {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	ret := new(model.TodoItem)
	row := conn.Db.QueryRowContext(ctx, query.TodoSelectItemById, id)
	err := row.Scan(&ret.Id, &ret.IdTodo, &ret.Title, &ret.Description, &ret.Flags,
		&ret.CreatedAt, &ret.UpdatedAt)
	return ret, err
}

func (t *Todo) SelectItemsByTodoId(ctx context.Context, id int32) ([]model.TodoItem, error) {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	rows, err := conn.Db.QueryContext(ctx, query.TodoSelectItemsByTodoId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []model.TodoItem
	for rows.Next() {
		var item model.TodoItem
		err = rows.Scan(&item.Id, &item.IdTodo, &item.Title, &item.Description, &item.Flags,
			&item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ret = append(ret, item)
	}

	return ret, nil
}

func (t *Todo) IsExists(ctx context.Context, label string, userId int32) (bool, error) {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	row := conn.Db.QueryRowContext(ctx, query.TodoIsExists, label, userId)
	return util.DbDataIsExists(row)
}

func (t *Todo) ItemIsExists(ctx context.Context, todoId int32, title string) (bool, error) {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	row := conn.Db.QueryRowContext(ctx, query.TodoItemIsExists, todoId, title)
	return util.DbDataIsExists(row)
}

func (t *Todo) Insert(ctx context.Context, todo *model.Todo, userId int32) error {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	now := util.Now()
	_, err := util.DbTryExec(ctx, conn.Db, query.TodoInsert, userId, todo.Label, now)
	return err
}

func (t *Todo) InsertItem(ctx context.Context, item *model.TodoItem) error {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	slcs := []any{
		item.IdTodo,
		item.Title,
		item.Description,
		item.Flags,
		util.Now(),
	}

	plc := util.DbSqlPlaceholder(len(slcs), 1)
	_, err := util.DbTryExec(ctx, conn.Db, query.TodoInsertItems+plc, slcs...)
	return err
}

func (t *Todo) InsertItems(ctx context.Context, items []model.TodoItem) error {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	now := util.Now()

	// 5 fields, each items
	slcs := make([]any, 0, len(items)*5)
	for i := range items {
		x := &items[i]
		slcs = append(slcs, x.IdTodo, x.Title, x.Description, x.Flags, now)
	}

	plc := util.DbSqlPlaceholder(5, len(items))
	_, err := util.DbTryExec(ctx, conn.Db, query.TodoInsertItems+plc, slcs...)
	return err
}

func (t *Todo) UpdateItemFlags(ctx context.Context, id, flags int32) error {
	conn := t.db.GetConn()
	defer t.db.PutConn(conn)

	now := util.Now()
	err := util.DbTxExecWithHandler(ctx, conn.Db, func(ctx context.Context, tx *sql.Tx, _ ...any) error {
		_, err := util.DbTxTryExecPartial(ctx, tx, query.TodoInsertItemHistory, now, id)
		if err != nil {
			return err
		}

		_, err = util.DbTxTryExecPartial(ctx, tx, query.TodoUpdateItemFlags, flags, now, id)
		return err
	})

	return err
}
