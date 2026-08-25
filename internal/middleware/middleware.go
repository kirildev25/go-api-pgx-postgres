package middleware

import (
	"context"
	"net/http"

	"github.com/kirildev25/go-api-pgx-postgres/internal/store"
)

type UserMiddleware struct {
	UserStore store.UserStore
}

type contextKey string

const userContextKey = contextKey("user")

func SetUser(r *http.Request, user *store.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func GetUser(r *http.Request) *store.User {
	user, ok := r.Context().Value(userContextKey).(*store.User)
	if !ok {
		return nil
	}
	return user
}
