package middleware

import (
	"log"
	"net/http"

	"github.com/rlapz/mmweb/service"
	"github.com/rlapz/mmweb/util"
)

const (
	FLAG_AUTH = (1 << 0)
)

type Middleware struct {
	http.ServeMux

	auth    util.Set
	items   []func(next http.Handler) http.Handler
	service *service.Service
}

func New(srv *service.Service) *Middleware {
	m := new(Middleware)

	m.service = srv
	m.auth = util.SetNew()
	m.addItems()

	return m
}

func (m *Middleware) AddHandler(path string, handler http.HandlerFunc, flags int) {
	if (flags & FLAG_AUTH) != 0 {
		m.auth.Add(path)
	}

	m.ServeMux.HandleFunc(path, handler)

	log.Println("path:", path)
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
