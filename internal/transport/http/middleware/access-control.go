package middleware

import (
	"net/http"
)

func AccessControl(next http.Handler) http.Handler {
	const fn = "internal.transport.http.AccessControl"
	h := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(h)
}
