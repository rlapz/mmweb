package middleware

import (
	"errors"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/util"
)

const bear = "Bearer "

func (m *Middleware) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.auth.Check(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		tokenlen := len(token)
		if tokenlen <= 0 {
			util.HttpErrBadRequest(w, errorx.AuthTokenNotFound.Error())
			return
		}

		bearLen := len(bear)
		if tokenlen <= bearLen {
			util.HttpErrBadRequest(w, errorx.AuthTokenInvalid.Error())
			return
		}

		if token[:bearLen] != bear {
			util.HttpErrBadRequest(w, errorx.AuthMethodInvalid.Error())
			return
		}

		claims, err := m.service.AuthVerify(r.Context(), token[bearLen:])
		switch {
		case err == nil: // OK
		case errors.Is(err, errorx.AuthTokenInvalid):
			util.HttpErrUnauthorized(w, err.Error())
			return
		default:
			util.HttpErrInternal(w, err, "failed to check token")
			return
		}

		util.ContextSetJwtClaims(&r, claims)

		next.ServeHTTP(w, r)
	})
}
