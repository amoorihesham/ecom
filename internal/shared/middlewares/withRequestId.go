package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type requestIdKey string

const ReqIdKey requestIdKey = "req-id"

func WithRequestId() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			reqId := uuid.NewString()

			ctx := context.WithValue(r.Context(), ReqIdKey, reqId)
			r = r.WithContext(ctx)
			next(w, r)
		})
	}
}
