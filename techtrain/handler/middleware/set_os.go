package middleware

import (
	"net/http"
	"techtrain-go-practice/context/os"

	"github.com/mileusna/useragent"
)

func SetOS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := os.ContextWithOS(r.Context(), ua.OS)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
