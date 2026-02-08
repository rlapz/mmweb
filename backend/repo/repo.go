package repo

import (
	"context"

	"github.com/rlapz/mmweb/model"
)

type Repo interface {
	UserSelectPasswordByName(ctx context.Context, uname string) (string, error)

	TodoInsert(ctx context.Context, uname string, todo *model.Todo) error
	TodoSelectById(ctx context.Context, id int32) (*model.Todo, error)
	TodoSelectByUsername(ctx context.Context, uname string) ([]model.Todo, error)

	AuthTokenInvalidInsert(ctx context.Context, token string) error
	AuthTokenInvalidCheck(ctx context.Context, token string) error
}
