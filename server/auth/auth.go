package auth

import (
	"context"
	"net/http"

	"github.com/go-pg/pg/v10"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type AuthUser struct {
	Name   string
	IsAuth bool
}

func Middleware(db *pg.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			user := &AuthUser{
				Name:   "Hello World",
				IsAuth: true,
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})

	}
}

func ForContext(ctx context.Context) *AuthUser {
	raw, _ := ctx.Value(userCtxKey).(*AuthUser)
	return raw
}
