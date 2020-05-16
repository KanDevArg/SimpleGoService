package middleware

import (
	"net/http"
	"time"

	log "github.com/op/go-logging"
	"golang.org/x/time/rate"
)

var (
	loggerMw = log.MustGetLogger("middleware")
)

// Middleware ...
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware ...
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		loggerMw.Infof("Request completed in %d ms", time.Since(start).Microseconds())
	})
}

// RateLimiter ...
func RateLimiter(limit int) Middleware {
	return func(next http.Handler) http.Handler {
		limiter := rate.NewLimiter(1, 200)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				w.WriteHeader(http.StatusTooManyRequests)

				e := http.StatusText(http.StatusTooManyRequests)
				w.Write([]byte(e))

				loggerMw.Warningf("Request from %s dropped (rate-limited)", r.RemoteAddr)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// CombineHandlers ...
func CombineHandlers(root http.Handler, mwf ...Middleware) (handler http.Handler) {
	handler = root
	for _, m := range mwf {
		handler = m(handler)
	}
	return
}
