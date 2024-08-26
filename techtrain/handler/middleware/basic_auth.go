package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func checkAuth(r *http.Request) bool {
	wantUserID := os.Getenv("BASIC_AUTH_USER_ID")
	wantPassword := os.Getenv("BASIC_AUTH_PASSWORD")

	userID, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return userID == wantUserID && password == wantPassword
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !checkAuth(r) {
			w.Header().Add("WWW-Authenticate", `Basic realm="secret"`)
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		_, err := fmt.Fprintf(w, "Successfully authenticated.\n")
		if err != nil {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r)
	})
}
