package auth

import (
	"context"
	"net/http"

	"github.com/SBPH-Matthew/testosterone-tracker/graph/model"
	"github.com/go-pg/pg/v10"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(db *pg.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			var myInt int32 = 24

			user := &model.User{
				FirstName: "Hello",
				LastName:  "World",
				Email:     "hello.world@gmail.com",
				ID:        "Testing",
				Gender:    "Male",
				Age:       &myInt,
				Password:  "testinggaming",
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
