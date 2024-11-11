package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func LoggingMiddleware(logger *logrus.Logger) func(next http.Handler) http.Handler {
	const fn = "internal.transport.http.middleware"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// format: method url remote_address duration
			start := time.Now()
			method := r.Method
			resource := r.URL.String()
			remoteAddr := r.RemoteAddr
			next.ServeHTTP(w, r)
			logger.Infof("[%v] %v, remote_addr=%v, duration=%s", method, resource, remoteAddr, time.Since(start).String())
		})
	}
}
