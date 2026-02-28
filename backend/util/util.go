package util

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/errorx"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotSlice   = errors.New("not a slice")
	ErrNotStruct  = errors.New("no a struct")
	ErrSliceZero  = errors.New("slice has zero item")
	ErrStructZero = errors.New("slice has zero field")
)

func JwtMakeSignedToken(method *jwt.SigningMethodHMAC, key []byte, issuer string,
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

func ContextSetJwtClaims(r **http.Request, cl jwt.MapClaims) {
	ctx := context.WithValue(context.Background(), config.CLAIMS_CONTEXT_NAME, cl)
	*r = (*r).WithContext(ctx)
}

func ContextGetJwtClaims(ctx context.Context) jwt.MapClaims {
	return ctx.Value(config.CLAIMS_CONTEXT_NAME).(jwt.MapClaims)
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
