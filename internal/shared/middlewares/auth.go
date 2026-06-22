package middleware

import (
	"ecom/internal/shared/httpx"
	"ecom/internal/shared/jwt"
	"net/http"
)

func Auth(jwtService jwt.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("Authorization")
			if token == "" {
				httpx.Error(w, 401, httpx.ErrUnauthorized, "missing token")
				return
			}

			user, err := jwtService.Parse(token)
			if err != nil {
				httpx.Error(w, 401, httpx.ErrUnauthorized, "invalid token")
				return
			}

			ctx := httpx.SetUser(r.Context(), *user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRoles(roles ...string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, ok := httpx.GetAuthUser(r.Context())
			if !ok {
				httpx.Error(w, 401, httpx.ErrUnauthorized, "unauthorized")
				return
			}

			for _, role := range roles {
				if user.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			httpx.Error(w, 403, httpx.ErrForbidden, "forbidden")
		})
	}
}
