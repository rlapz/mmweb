package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/util"
)

const bear = "Bearer"

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

		token := strings.ReplaceAll(auth, "Bearer ", "")

		isOk, err := m.service.AuthTokenCheck(r.Context(), token)
		if err != nil {
			util.HttpErrInternal(w, err, "failed to check token")
			return
		}

		if !isOk {
			util.HttpErrUnauthorized(w, "invalid token")
			return
		}

		tokenSigned, err := jwt.Parse(token, func(tok *jwt.Token) (any, error) {
			mth, ok := tok.Method.(*jwt.SigningMethodHMAC)
			if !ok || mth != m.signMethod {
				return nil, errorx.AuthSignMethod
			}

			return m.signKey, nil
		})

		if err != nil {
			util.HttpErrUnauthorized(w, err.Error())
			return
		}

		claims, ok := tokenSigned.Claims.(jwt.MapClaims)
		if !ok || !tokenSigned.Valid {
			util.HttpErrUnauthorized(w, errorx.AuthTokenClaims.Error())
			return
		}

		util.ContextSetJwtClaims(&r, claims)

		next.ServeHTTP(w, r)
	})
}
