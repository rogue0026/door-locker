package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type responseData struct {
	size   int
	status int
}

type Wrapper struct {
	http.ResponseWriter
	rd *responseData
}

func (w *Wrapper) Write(p []byte) (int, error) {
	sz, err := w.ResponseWriter.Write(p)
	w.rd.size += sz
	return sz, err
}

func (w *Wrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.rd.status = statusCode
}

func LoggingMiddleware(logger *logrus.Logger) func(next http.Handler) http.Handler {
	const fn = "internal.transport.http.LoggingMiddleware"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			method := r.Method
			resource := r.URL.String()
			remoteAddr := r.RemoteAddr
			wrapper := Wrapper{
				ResponseWriter: w,
				rd:             new(responseData),
			}
			next.ServeHTTP(&wrapper, r)
			logger.Infof("[%v] %v, status_code=%d, remote_addr=%s, duration=%s, output_data_len=%d bytes",
				method,
				resource,

				wrapper.rd.status,
				remoteAddr,
				time.Since(start).String(),
				wrapper.rd.size)
		})
	}
}
