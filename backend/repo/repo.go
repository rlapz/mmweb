package repo

import (
	"context"

	"github.com/rlapz/mmweb/model"
)

// TODO: add pagination for multiple items
type Auth interface {
	SelectByToken(ctx context.Context, token string) (*model.Auth, error)
	Insert(ctx context.Context, auth *model.Auth) error
	UpdateFlagsByToken(ctx context.Context, token string, flags int32) error
}

type Todo interface {
	SelectById(ctx context.Context, id int32) (*model.Todo, error)
	SelectByUserId(ctx context.Context, id int32) ([]model.Todo, error)
	SelectItemById(ctx context.Context, id int32) (*model.TodoItem, error)
	SelectItemsByTodoId(ctx context.Context, id int32) ([]model.TodoItem, error)

	IsExists(ctx context.Context, label string, userId int32) (bool, error)
	IsExistsById(ctx context.Context, id int32) (bool, error)
	ItemIsExists(ctx context.Context, todoId int32, title string) (bool, error)

	Insert(ctx context.Context, todo *model.Todo) error
	InsertItem(ctx context.Context, items *model.TodoItem) error
	InsertItems(ctx context.Context, items []model.TodoItem) error

	Update(ctx context.Context, todo *model.Todo) error
	UpdateIsActive(ctx context.Context, id int32, isActive bool) error
	UpdateItem(ctx context.Context, item *model.TodoItem) error
	UpdateItemStatus(ctx context.Context, id, status int32) error
}

type User interface {
	SelectByName(ctx context.Context, uname string) (*model.User, error)
	Insert(ctx context.Context, user *model.User) error
	IsExists(ctx context.Context, uname string) (bool, error)
}
