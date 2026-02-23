package middleware

import (
	"net/http"
	"strings"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/util"
)

const bear = "Bearer "

func (m *Middleware) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.authExcluded.Check(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		auth := r.Header.Get("Authorization")
		if !strings.Contains(auth, bear) {
			util.HttpErrBadRequest(w, errorx.AuthTokenNotFound.Error())
			return
		}

		token := strings.ReplaceAll(auth, bear, "")

		isOk, err := m.service.AuthVerify(r.Context(), token)
		if err != nil {
			util.HttpErrInternal(w, err, "failed to check token")
			return
		}

		if !isOk {
			util.HttpErrUnauthorized(w, "invalid token")
			return
		}

		claims, err := m.service.AuthToken(token)
		if err != nil {
			util.HttpErrUnauthorized(w, err.Error())
			return
		}

		util.ContextSetJwtClaims(&r, claims)

		next.ServeHTTP(w, r)
	})
}
