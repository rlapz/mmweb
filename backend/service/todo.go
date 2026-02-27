package service

import (
	"context"

	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

func (s *Service) TodoAdd(ctx context.Context, todo *model.Todo, uname string) error {
	err := util.ValidateStruct(ctx, todo)
	if err != nil {
		return err
	}

	id, err := s.repoUser.SelectIdByName(ctx, uname)
	if err != nil {
		return err
	}

	return s.repoTodo.Insert(ctx, todo, id)
}

func (s *Service) TodoAddItems(ctx context.Context, id int32, items []model.TodoItem) error {
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

	return s.repoTodo.InsertItems(ctx, items)
}

func (s *Service) TodoGet(ctx context.Context, id int32) (*model.Todo, error) {
	return s.repoTodo.SelectById(ctx, id)
}

func (s *Service) TodoGetList(ctx context.Context, uname string) ([]model.Todo, error) {
	id, err := s.repoUser.SelectIdByName(ctx, uname)
	if err != nil {
		return nil, err
	}

	return s.repoTodo.SelectByUserId(ctx, id)
}

func (s *Service) TodoGetItems(ctx context.Context, id int32) ([]model.TodoItem, error) {
	return s.repoTodo.SelectItemsById(ctx, id)
}
