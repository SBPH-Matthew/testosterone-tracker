package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/SBPH-Matthew/testosterone-tracker/dbmodels"
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

			var tokenString string

			// 1. Check Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				fmt.Println("Found Authorization header:", authHeader)
				if strings.HasPrefix(authHeader, "Bearer ") {
					tokenString = strings.TrimPrefix(authHeader, "Bearer ")
				}
			}

			// 2. If no header, try cookie
			if tokenString == "" {
				c, err := r.Cookie("auth-cookie")
				if err == nil {
					tokenString = c.Value
				}
			}

			// If still no token → user stays unauthenticated
			if tokenString == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate token
			token, err := services.ValidateToken(tokenString)
			if err != nil || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

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

			// Fetch user
			user := &dbmodels.User{}
			err = db.Model(user).Where("id = ?", userID).First()
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			graphQlUser := &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Password:  user.Password,
				Gender:    user.Gender,
				Age:       user.Age,
				Token:     &tokenString,
			}

			ctx := context.WithValue(r.Context(), userCtxKey, graphQlUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
