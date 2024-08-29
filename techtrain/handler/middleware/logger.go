package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"techtrain-go-practice/context/os"
	"time"
)

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTime := time.Now()

		next.ServeHTTP(w, r)

		latency := time.Since(accessTime).Milliseconds()
		path := r.URL.Path
		os, err := os.OSFromContext(r.Context())
		if err != nil {
			log.Println(err)
			return
		}

		accessLog := Log{
			Timestamp: accessTime,
			Latency:   latency,
			Path:      path,
			OS:        os,
		}

		bytes, err := json.Marshal(accessLog)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(bytes))
	})
}
