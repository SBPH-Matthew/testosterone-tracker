package auth

import (
	"context"
	"net/http"

	"github.com/SBPH-Matthew/testosterone-tracker/graph/model"
	"github.com/SBPH-Matthew/testosterone-tracker/services"
	"github.com/go-pg/pg/v10"
	"github.com/golang-jwt/jwt/v5"
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

func MiddlewareV2(db *pg.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from cookie

			c, err := r.Cookie("auth-cookie")
			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			// Validate token
			token, err := services.ValidateToken(c.Value)
			if err != nil || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			// Extract user id from claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			userID, ok := claims["user_id"].(string)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			// Fetch user from database
			user := &model.User{}
			err = db.Model(user).Where("id = ?", userID).First()
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
