package util

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rlapz/mmweb/config"
	"golang.org/x/crypto/bcrypt"
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

func ContextSetJwtClaims(r **http.Request, cl jwt.Claims) {
	ctx := context.WithValue(context.Background(), config.CLAIMS_CONTEXT_NAME, cl)
	*r = (*r).WithContext(ctx)
}

func ContextGetJwtClaims(ctx context.Context) jwt.MapClaims {
	return ctx.Value(config.CLAIMS_CONTEXT_NAME).(jwt.MapClaims)
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

func Cond[T any](cnd bool, expected, alt T) T {
	if cnd {
		return expected
	}

	return alt
}
