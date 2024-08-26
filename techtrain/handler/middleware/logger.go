package middleware

import (
	"log"
	"net/http"
	"techtrain-go-practice/context/os"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTime := time.Now()

		next.ServeHTTP(w, r)

		latency := time.Since(accessTime)
		path := r.URL.Path
		os, err := os.OSFromContext(r.Context())
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf(`{"timestamp": "%s", "latency": "%s", "path": "%s", "os": "%s"}`, time.Now(), latency, path, os)
	})
}
