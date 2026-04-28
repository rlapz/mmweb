package middleware

import (
	"errors"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model/api"
)

const _bear = "Bearer "
const _bearLen = len(_bear)

func (m *Middleware) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.auth.Check(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		tokenlen := len(token)
		if tokenlen <= 0 {
			api.HttpErrBadRequest(w, errorx.AuthTokenNotFound.Error())
			return
		}

		if tokenlen <= _bearLen {
			api.HttpErrBadRequest(w, errorx.AuthTokenInvalid.Error())
			return
		}

		if token[:_bearLen] != _bear {
			api.HttpErrBadRequest(w, errorx.AuthMethodInvalid.Error())
			return
		}

		ctx, err := m.service.AuthContextNew(r.Context(), token[_bearLen:])
		switch {
		case err == nil: // OK
		case errors.Is(err, errorx.AuthTokenInvalid):
			api.HttpErrUnauthorized(w, err.Error())
			return
		default:
			api.HttpErrInternal(w, err, "failed to check token")
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
