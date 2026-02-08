package service

import (
	"context"
	"database/sql"
	"errors"
)

func (s *Service) AuthTokenCheck(ctx context.Context, token string) (bool, error) {
	err := s.repo.AuthTokenInvalidCheck(ctx, token)
	if errors.Is(err, sql.ErrNoRows) {
		return true, nil
	}

	return false, err
}
