package service

import (
	"context"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
)

func (s *Service) AddTodo(ctx context.Context, uname string, todo *model.Todo) error {
	if todo.Title == "" {
		return errorx.DataInvalid
	}

	return s.repo.TodoInsert(ctx, uname, todo)
}

func (s *Service) GetTodoById(ctx context.Context, id int32) (*model.Todo, error) {
	return s.repo.TodoSelectById(ctx, id)
}

func (s *Service) GetTodoByUsername(ctx context.Context, uname string) ([]model.Todo, error) {
	return s.repo.TodoSelectByUsername(ctx, uname)
}
