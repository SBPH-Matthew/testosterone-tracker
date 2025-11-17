package auth

import (
	"net/http"

	"github.com/go-pg/pg/v10"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type User struct {
	Email string
}

func Middleware(db *pg.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}
		})
	}
}
