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

func (s *Service) TodoAdd(ctx context.Context, todo *model.Todo) error {
	err := util.ValidateStruct(ctx, todo)
	if err != nil {
		return err
	}

	todo.IsActive = true
	todo.Label = strings.ToLower(todo.Label)
	isExists, err := s.repoTodo.IsExistsByLabel(ctx, todo.Label, todo.IdUser)
	if err != nil {
		return err
	}

	if isExists {
		return errorx.DataExists
	}

	return s.repoTodo.Insert(ctx, todo)
}

func (s *Service) TodoAddItems(ctx context.Context, id int32, items []model.TodoItem) error {
	isExists, err := s.repoTodo.IsExists(ctx, id)
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

	isExists, err := s.repoTodo.IsExists(ctx, item.IdTodo)
	if err != nil {
		return err
	}

	if !isExists {
		return errorx.NoDataFound
	}

	item.Title = strings.ToLower(item.Title)
	item.Description = strings.ToLower(item.Description)
	isExists, err = s.repoTodo.ItemIsExistsByTitle(ctx, item.IdTodo, item.Title)
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

func (s *Service) TodoGetList(ctx context.Context, userId int32, pag *model.Pagination) ([]model.Todo, error) {
	return s.repoTodo.SelectByUserId(ctx, userId, pag)
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

func (s *Service) TodoUpdate(ctx context.Context, todo *model.Todo) error {
	todo.Label = strings.ToLower(todo.Label)
	isExists, err := s.repoTodo.IsExistsByLabel(ctx, todo.Label, todo.IdUser)
	if err != nil {
		return err
	}

	if isExists {
		return errorx.DataExists
	}

	return s.repoTodo.Update(ctx, todo)
}

func (s *Service) TodoDelete(ctx context.Context, id int32) error {
	isExists, err := s.repoTodo.IsExists(ctx, id)
	if err != nil {
		return err
	}

	if !isExists {
		return errorx.NoDataFound
	}

	return s.repoTodo.UpdateIsActive(ctx, id, false)
}

func (s *Service) TodoUpdateItem(ctx context.Context, item *model.TodoItem) error {
	err := util.ValidateStruct(ctx, item)
	if err != nil {
		return err
	}

	item.Title = strings.ToLower(item.Title)
	item.Description = strings.ToLower(item.Description)
	return s.repoTodo.UpdateItem(ctx, item)
}
