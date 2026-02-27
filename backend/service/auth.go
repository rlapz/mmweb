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

func (s *Service) AuthVerify(ctx context.Context, token string) (jwt.MapClaims, error) {
	claims, err := s.authParseToken(token)
	if err != nil {
		return nil, err
	}

	flags, err := s.repo.AuthTokenSelectFlags(ctx, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	if flags != model.AUTH_FLAG_LOGGED_IN {
		return nil, errorx.AuthTokenInvalid
	}

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
	err := util.ValidateStruct(ctx, user)
	if err != nil {
		return err
	}

	isExists, err := s.repo.UserIsExists(ctx, user.Name)
	if err != nil {
		return err
	}

	if isExists {
		return errorx.DataExists
	}

	user.Password, err = util.HashPasswordGenerate(user.Password)
	if err != nil {
		return err
	}

	return s.repo.UserInsert(ctx, user)
}

// Private
func (s *Service) authParseToken(token string) (jwt.MapClaims, error) {
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
