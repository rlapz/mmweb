package repo

import (
	"context"

	"github.com/rlapz/mmweb/model"
)

type Repo interface {
	UserSelectIdByName(ctx context.Context, uname string) (int32, error)
	UserSelectPasswordByName(ctx context.Context, uname string) (string, error)
	UserInsert(ctx context.Context, user *model.User) error
	UserIsExists(ctx context.Context, uname string) (bool, error)

	TodoInsert(ctx context.Context, todo *model.Todo, userId int32) error
	TodoInsertItems(ctx context.Context, items []model.TodoItem) error
	TodoIsExists(ctx context.Context, label string, userId int32) (bool, error)
	TodoSelectById(ctx context.Context, id int32) (*model.Todo, error)
	TodoSelectByUserId(ctx context.Context, id int32) ([]model.Todo, error)
	TodoSelectItemsById(ctx context.Context, id int32) ([]model.TodoItem, error)

	AuthTokenInsert(ctx context.Context, token string, flags int) error
	AuthTokenUpdateFlags(ctx context.Context, token string, flags int) error
	AuthTokenSelectFlags(ctx context.Context, token string) (int, error)
}
