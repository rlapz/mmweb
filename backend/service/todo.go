package service

import (
	"context"

	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

func (s *Service) AddTodo(ctx context.Context, todo *model.Todo, uname string) error {
	err := util.ValidateStruct(ctx, todo)
	if err != nil {
		return err
	}

	id, err := s.repo.UserSelectIdByName(ctx, uname)
	if err != nil {
		return err
	}

	return s.repo.TodoInsert(ctx, todo, id)
}

func (s *Service) AddTodoItems(ctx context.Context, id int32, items []model.TodoItem) error {
	for i := range items {
		item := &items[i]
		err := util.ValidateStruct(ctx, item)
		if err != nil {
			return err
		}
	}

	for i := range items {
		v := &items[i]
		v.IdTodo = id
	}

	return s.repo.TodoInsertItems(ctx, items)
}

func (s *Service) GetTodo(ctx context.Context, id int32) (*model.Todo, error) {
	return s.repo.TodoSelectById(ctx, id)
}

func (s *Service) GetTodoList(ctx context.Context, uname string) ([]model.Todo, error) {
	id, err := s.repo.UserSelectIdByName(ctx, uname)
	if err != nil {
		return nil, err
	}

	return s.repo.TodoSelectByUserId(ctx, id)
}

func (s *Service) GetTodoItems(ctx context.Context, id int32) ([]model.TodoItem, error) {
	return s.repo.TodoSelectItemsById(ctx, id)
}
