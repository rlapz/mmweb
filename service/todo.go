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

	return s.repo.InsertTodo(ctx, uname, todo)
}
