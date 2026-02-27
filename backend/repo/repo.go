package repo

import (
	"context"

	"github.com/rlapz/mmweb/model"
)

type Auth interface {
	TokenInsert(ctx context.Context, token string, flags int) error
	TokenUpdateFlags(ctx context.Context, token string, flags int) error
	TokenSelectFlags(ctx context.Context, token string) (int, error)
}

type Todo interface {
	Insert(ctx context.Context, todo *model.Todo, userId int32) error
	InsertItems(ctx context.Context, items []model.TodoItem) error
	IsExists(ctx context.Context, label string, userId int32) (bool, error)
	SelectById(ctx context.Context, id int32) (*model.Todo, error)
	SelectByUserId(ctx context.Context, id int32) ([]model.Todo, error)
	SelectItemsById(ctx context.Context, id int32) ([]model.TodoItem, error)
}

type User interface {
	SelectIdByName(ctx context.Context, uname string) (int32, error)
	SelectPasswordByName(ctx context.Context, uname string) (string, error)
	Insert(ctx context.Context, user *model.User) error
	IsExists(ctx context.Context, uname string) (bool, error)
}
