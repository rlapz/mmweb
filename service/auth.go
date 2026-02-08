package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/util"
)

func (s *Service) AuthUser(ctx context.Context, uname, passwd string) error {
	hashed, err := s.repo.UserSelectPasswordByName(ctx, uname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorx.AuthInvalidCredential
		}

		return err
	}

	if !util.HashPasswordVerify(hashed, passwd) {
		return errorx.AuthInvalidCredential
	}

	return nil
}

func (s *Service) AuthTokenCheck(ctx context.Context, token string) (bool, error) {
	err := s.repo.AuthTokenInvalidCheck(ctx, token)
	if errors.Is(err, sql.ErrNoRows) {
		return true, nil
	}

	return false, err
}

func (s *Service) AuthTokenAdd(ctx context.Context, token string) error {
	return s.repo.AuthTokenInvalidInsert(ctx, token)
}
