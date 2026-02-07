package middleware

import (
	"log"
	"net/http"
)

func (m *Middleware) LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

		log.Printf("%s: %s: %+v: %+v\n", r.RemoteAddr, r.Method, r.RequestURI, r.Header)
	})
}
