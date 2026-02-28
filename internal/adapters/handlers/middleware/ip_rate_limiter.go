package middleware

import (
	"net"
	"net/http"
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
)

func NewIPRateLimiter(limiter ports.RateLimiter, window time.Duration, limit int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr // fallback
			}
			allowed, _ := limiter.Allow(r.Context(), ip, window, limit)
			if !allowed {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
