package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/util"
)

func (s *Service) Auth(ctx context.Context, uname, passwd string) error {
	hashed, err := s.repo.SelectUserPasswordByName(ctx, uname)
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
