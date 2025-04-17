package app

import (
	"context"
	"net/http"
	"strings"
)

func (s *server) CheckJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		token = strings.TrimPrefix(token, "Bearer ")

		var ctx context.Context

		claims, err := s.userService.ValidateToken(token)

		if err != nil {
			ctx = context.WithValue(r.Context(), "validate-error", err.Error())
		} else {
			ctx = context.WithValue(r.Context(), "role", claims.Role)
			ctx = context.WithValue(ctx, "user_id", claims.Id)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
