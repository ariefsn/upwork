package middlewares

import (
	"context"
	"net/http"

	"github.com/ariefsn/upwork/constant"
	"github.com/ariefsn/upwork/env"
)

func Inject(env env.Env) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			// Inject Writer
			ctx := context.WithValue(r.Context(), constant.WriterCtxKey, w)
			r = r.WithContext(ctx)

			// Inject Request
			ctx = context.WithValue(ctx, constant.HttpRequestCtxKey, r)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
