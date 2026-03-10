package util

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rlapz/mmweb/errorx"
	"golang.org/x/crypto/bcrypt"
)

func JwtNewToken(method *jwt.SigningMethodHMAC, key []byte, issuer string,
	exp time.Duration) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"iss": issuer,
		"jti": uuid.NewString(),
		"iat": now.Unix(),
		"exp": now.Add(exp).Unix(),
	})

	return token.SignedString(key)
}

func ContextSleep(ctx context.Context, dur time.Duration) error {
	select {
	case <-time.After(dur):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func HashPasswordGenerate(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(b), err
}

func HashPasswordVerify(hash, plain string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil {
		return true
	}

	return false
}

func Now() int64 {
	return time.Now().Unix()
}

func Cnd[T any](cond bool, expected, alt T) T {
	if cond {
		return expected
	}

	return alt
}

func Cnd2[T any](cond bool, expected T, ret *T) {
	if cond {
		*ret = expected
	}
}

func ValidateStruct(ctx context.Context, item any) error {
	vl := validator.New()
	err := vl.StructCtx(ctx, item)
	if err != nil {
		return errors.Join(err, errorx.DataInvalid)
	}

	return nil
}

func ParseJsonReader[T any](reader io.Reader) (*T, error) {
	ret := new(T)
	if err := json.NewDecoder(reader).Decode(ret); err != nil {
		return nil, err
	}

	return ret, nil
}
