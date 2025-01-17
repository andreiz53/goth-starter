package server

import (
	"context"
	"goth/types"
	"net/http"
	"strings"
)

func (s Server) WithUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/assets") {
			next.ServeHTTP(w, r)
			return
		}

		u := s.getAuthenticatedUser(r)
		if !u.LoggedIn {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), types.UserContextKey, u)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
