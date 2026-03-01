package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/rlapz/mmweb/errorx"
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

	todo.IsActive = true
	todo.Label = strings.ToLower(todo.Label)
	isExists, err := s.repoTodo.IsExists(ctx, todo.Label, id)
	if err != nil {
		return err
	}

	if isExists {
		return errorx.DataExists
	}

	return s.repoTodo.Insert(ctx, todo, id)
}

func (s *Service) TodoAddItems(ctx context.Context, id int32, items []model.TodoItem) error {
	isExists, err := s.repoTodo.IsExistsById(ctx, id)
	if err != nil {
		return err
	}

	if !isExists {
		return errorx.NoDataFound
	}

	for i := range items {
		item := &items[i]
		item.IdTodo = id

		err := util.ValidateStruct(ctx, item)
		if err != nil {
			return err
		}

		item.Title = strings.ToLower(item.Title)
		item.Description = strings.ToLower(item.Description)
	}

	return s.repoTodo.InsertItems(ctx, items)
}

func (s *Service) TodoAddItem(ctx context.Context, item *model.TodoItem) error {
	err := util.ValidateStruct(ctx, item)
	if err != nil {
		return err
	}

	isExists, err := s.repoTodo.IsExistsById(ctx, item.IdTodo)
	if err != nil {
		return err
	}

	if !isExists {
		return errorx.NoDataFound
	}

	item.Title = strings.ToLower(item.Title)
	item.Description = strings.ToLower(item.Description)
	isExists, err = s.repoTodo.ItemIsExists(ctx, item.IdTodo, item.Title)
	if err != nil {
		return err
	}

	if isExists {
		return errorx.DataExists
	}

	return s.repoTodo.InsertItem(ctx, item)
}

func (s *Service) TodoGet(ctx context.Context, id int32) (*model.Todo, error) {
	ret, err := s.repoTodo.SelectById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NoDataFound
		}

		return nil, err
	}

	return ret, nil
}

func (s *Service) TodoGetList(ctx context.Context, uname string) ([]model.Todo, error) {
	id, err := s.repoUser.SelectIdByName(ctx, uname)
	if err != nil {
		return nil, err
	}

	return s.repoTodo.SelectByUserId(ctx, id)
}

func (s *Service) TodoGetItem(ctx context.Context, id int32) (*model.TodoItem, error) {
	ret, err := s.repoTodo.SelectItemById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NoDataFound
		}

		return nil, err
	}

	return ret, nil
}

func (s *Service) TodoGetItemList(ctx context.Context, id int32) ([]model.TodoItem, error) {
	return s.repoTodo.SelectItemsByTodoId(ctx, id)
}
