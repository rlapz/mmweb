package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/src/errorx"
	"github.com/rlapz/mmweb/src/util"
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

		tokr := strings.ReplaceAll(auth, "Bearer ", "")
		tok, err := jwt.Parse(tokr, func(tok *jwt.Token) (any, error) {
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

		cl, ok := tok.Claims.(jwt.MapClaims)
		if !ok || !tok.Valid {
			util.HttpErrUnauthorized(w, errorx.AuthTokenClaims.Error())
			return
		}

		r = util.JwtClaimsSetContext(r, cl)

		next.ServeHTTP(w, r)
	})
}
