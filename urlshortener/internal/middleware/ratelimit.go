package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
)

// limit up to 2 requests per second, with a burst allowance of 5 requests.
var limiter = rate.NewLimiter(2, 5)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
