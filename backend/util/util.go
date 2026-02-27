package util

import (
	"context"
	"errors"
	"net/http"
	"reflect"
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

func UnwrapPointer(item any) reflect.Value {
	typ := reflect.ValueOf(item)
	for typ.Kind() == reflect.Pointer {
		typ = reflect.Indirect(typ)
	}

	return typ
}

func UnwrapPointerType(item any) reflect.Type {
	typ := reflect.TypeOf(item)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	return typ
}

func StructFieldsCount(item any) int {
	return UnwrapPointerType(item).NumField()
}

// TODO: convert all data types and support variadic parameter
func StructToAnySlice(items any) ([]any, error) {
	typ := UnwrapPointer(items)
	if typ.Kind() == reflect.Struct {
		return structToAnySliceBuilder(typ)
	}

	switch typ.Kind() {
	case reflect.Array, reflect.Slice:
	default:
		return nil, ErrNotSlice
	}

	count := typ.Len()
	if count == 0 {
		return nil, ErrSliceZero
	}

	item := typ.Index(0)
	if item.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	slcs := make([]any, 0, (item.NumField() * count))
	for i := range count {
		slice, err := structToAnySliceBuilder(UnwrapPointer(typ.Index(i).Interface()))
		if err != nil {
			return nil, err
		}

		slcs = append(slcs, slice...)
	}

	return slcs, nil
}

func ValidateStruct(ctx context.Context, item any) error {
	vl := validator.New()
	err := vl.StructCtx(ctx, item)
	if err != nil {
		return errors.Join(err, errorx.DataInvalid)
	}

	return nil
}

// Private
func structToAnySliceBuilder(typ reflect.Value) ([]any, error) {
	count := typ.NumField()
	ret := make([]any, count)
	for i := range count {
		ret[i] = typ.Field(i).Interface()
	}

	return ret, nil
}
