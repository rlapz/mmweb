package repo

import (
	"context"

	"github.com/rlapz/mmweb/model"
)

type Repo interface {
	SelectUserPasswordByName(ctx context.Context, uname string) (string, error)

	InsertTodo(ctx context.Context, uname string, todo *model.Todo) error
	SelectTodoById(ctx context.Context, id int32) (*model.Todo, error)
}
