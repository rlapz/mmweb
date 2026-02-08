package sqlite

import (
	"context"
	"time"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

func (r *Repo) TodoInsert(ctx context.Context, uname string, todo *model.Todo) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	todo.CreatedAt = time.Now()

	aff, err := util.DbTransactionTryExec(ctx, conn.Db, query.InsertTodo, todo.Title, todo.Description,
		todo.Flags, todo.CreatedAt, uname)
	if err != nil {
		return err
	}

	if aff == 0 {
		return errorx.NoDataSaved
	}

	return nil
}

func (r *Repo) TodoSelectById(ctx context.Context, id int32) (*model.Todo, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	ret := new(model.Todo)
	row := conn.Db.QueryRowContext(ctx, query.SelectTodoById, id)
	err := row.Scan(&ret.Id, &ret.IdUser, &ret.Title, &ret.Description, &ret.Flags,
		&ret.CreatedAt, &ret.CreatedBy)
	return ret, err
}
