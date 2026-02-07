package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/src/config"
	"github.com/rlapz/mmweb/src/util"
)

type Middleware struct {
	http.ServeMux

	signMethod   *jwt.SigningMethodHMAC
	signKey      []byte
	authExcluded util.Set
	items        []func(next http.Handler) http.Handler
}

func New(cfg *config.Config) *Middleware {
	m := new(Middleware)

	m.signMethod = cfg.JwtSignMethod
	m.signKey = cfg.JwtSignatureKey
	m.addItems()
	m.addAuthExcluded()

	return m
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var st http.Handler = &m.ServeMux
	for _, next := range m.items {
		st = next(st)
	}

	st.ServeHTTP(w, r)
}

func (m *Middleware) addItems() {
	m.items = []func(http.Handler) http.Handler{
		m.AuthHandler,
		m.LogHandler,
	}
}

func (m *Middleware) addAuthExcluded() {
	m.authExcluded = util.SetNew()
	m.authExcluded.Add([]string{
		"/",
		"/login",
	})
}
