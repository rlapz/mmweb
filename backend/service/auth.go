package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

func (s *Service) AuthToken(token string) (jwt.MapClaims, error) {
	tokSigned, err := jwt.Parse(token, func(tok *jwt.Token) (any, error) {
		mth, ok := tok.Method.(*jwt.SigningMethodHMAC)
		if !ok || mth != s.signMethod {
			return nil, errorx.AuthSignMethod
		}

		return s.signKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokSigned.Claims.(jwt.MapClaims)
	if !ok || !tokSigned.Valid {
		return nil, errorx.AuthTokenClaims
	}

	claims["token"] = token
	return claims, nil
}

func (s *Service) AuthLogin(ctx context.Context, uname, passwd string) (string, error) {
	hashed, err := s.repo.UserSelectPasswordByName(ctx, uname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errorx.AuthInvalidCredential
		}

		return "", err
	}

	if !util.HashPasswordVerify(hashed, passwd) {
		return "", errorx.AuthInvalidCredential
	}

	token, err := util.JwtMakeSignedToken(s.signMethod, s.signKey, uname, s.loginExp)
	if err != nil {
		return "", err
	}

	err = s.repo.AuthTokenInsert(ctx, token, model.AUTH_FLAG_LOGGED_IN)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) AuthLogout(ctx context.Context, token string) error {
	return s.repo.AuthTokenUpdateFlags(ctx, token, model.AUTH_FLAG_LOGGED_OUT)
}

func (s *Service) AuthRegister(ctx context.Context, user *model.User) error {
	// TODO: validate fields

	exists, err := s.repo.UserIsExists(ctx, user.Name)
	if err != nil {
		return err
	}

	if exists {
		return errorx.DataExists
	}

	user.Password, err = util.HashPasswordGenerate(user.Password)
	if err != nil {
		return err
	}

	return s.repo.UserInsert(ctx, user)
}

func (s *Service) AuthVerify(ctx context.Context, token string) (bool, error) {
	flags, err := s.repo.AuthTokenSelectFlags(ctx, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	if flags != model.AUTH_FLAG_LOGGED_IN {
		return false, nil
	}

	return true, nil
}
