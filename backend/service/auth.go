package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

func (s *Service) AuthContextNew(ctx context.Context, token string) (context.Context, error) {
	claims, err := s.authParseToken(token)
	if err != nil {
		return nil, errors.Join(errorx.AuthTokenInvalid, err)
	}

	auth, err := s.repoAuth.SelectByToken(ctx, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.AuthTokenInvalid
		}

		return nil, err
	}

	if auth.Flags != model.AUTH_FLAG_LOGGED_IN {
		return nil, errorx.AuthTokenInvalid
	}

	claims[config.CLAIMS_USER_ID_NAME] = auth.IdUser
	return context.WithValue(context.Background(), config.CLAIMS_CONTEXT_NAME, claims), nil
}

func (s *Service) AuthContextGetClaims(ctx context.Context) jwt.MapClaims {
	return ctx.Value(config.CLAIMS_CONTEXT_NAME).(jwt.MapClaims)
}

func (s *Service) AuthContextGetToken(ctx context.Context) string {
	return s.AuthContextGetClaims(ctx)[config.CLAIMS_TOKEN_NAME].(string)
}

func (s *Service) AuthContextGetUserId(ctx context.Context) int32 {
	return s.AuthContextGetClaims(ctx)[config.CLAIMS_USER_ID_NAME].(int32)
}

func (s *Service) AuthLogin(ctx context.Context, uname, passwd string) (string, error) {
	user, err := s.repoUser.SelectByName(ctx, uname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errorx.AuthInvalidCredential
		}

		return "", err
	}

	if !util.HashPasswordVerify(user.Password, passwd) {
		return "", errorx.AuthInvalidCredential
	}

	token, err := util.JwtNewToken(s.signMethod, s.signKey, user.Name, s.loginExp)
	if err != nil {
		return "", err
	}

	auth := model.Auth{
		IdUser: user.Id,
		Token:  token,
		Flags:  model.AUTH_FLAG_LOGGED_IN,
	}

	err = s.repoAuth.Insert(ctx, &auth)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) AuthLogout(ctx context.Context, token string) error {
	return s.repoAuth.UpdateFlagsByToken(ctx, token, model.AUTH_FLAG_LOGGED_OUT)
}

func (s *Service) AuthRegister(ctx context.Context, user *model.User) error {
	err := util.ValidateStruct(ctx, user)
	if err != nil {
		return err
	}

	isExists, err := s.repoUser.IsExistsByName(ctx, user.Name)
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

	return s.repoUser.Insert(ctx, user)
}

/**************
 * Private    *
 **************/
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

	claims[config.CLAIMS_TOKEN_NAME] = token
	return claims, nil
}
