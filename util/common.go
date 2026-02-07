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

type Set struct {
	item map[string]struct{}
}

func SetNew() Set {
	return Set{
		item: make(map[string]struct{}),
	}
}

func (s *Set) Add(entry []string) {
	for _, x := range entry {
		s.item[x] = struct{}{}
	}
}

func (s *Set) Del(entry string) {
	delete(s.item, entry)
}

func (s *Set) Check(entry string) bool {
	_, isOk := s.item[entry]
	return isOk
}

func JwtMakeClaims(method *jwt.SigningMethodHMAC, issuer string, exp time.Duration) *jwt.Token {
	now := time.Now()
	return jwt.NewWithClaims(method, jwt.MapClaims{
		"iss": issuer,
		"jti": uuid.NewString(),
		"iat": now.Unix(),
		"exp": now.Add(exp).Unix(),
	})
}

func JwtClaimsSetContext(r *http.Request, cl jwt.Claims) *http.Request {
	ctx := context.WithValue(context.Background(), config.CLAIMS_CONTEXT_NAME, cl)
	return r.WithContext(ctx)
}

func JwtClaimsGetContext(ctx context.Context) jwt.MapClaims {
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
