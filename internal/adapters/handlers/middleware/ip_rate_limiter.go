package middleware

import (
	"net"
	"net/http"
	"time"

	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/domain"
	"gitlab.com/yerdaulet.zhumabay/golang-hexagonal-architecture-template/internal/core/ports"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewIPRateLimiter(logger ports.Logger, limiter ports.RateLimiter, window time.Duration, limit int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			span := trace.SpanFromContext(ctx)

			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr
			}

			// Add the IP to the trace metadata for easier searching in Tempo
			span.SetAttributes(attribute.String("http.client_ip", ip))

			allowed, err := limiter.Allow(ctx, ip, window, limit)
			if err != nil {
				logger.Error(ctx, domain.LogLevelMiddleware, "Redis IP limiter error:", err)
			}

			if !allowed {
				// Mark the span as "Rate Limited"
				span.SetAttributes(attribute.Bool("http.rate_limited", true))

				logger.Info(ctx, domain.LogLevelMiddleware, "Request blocked by rate limiter", "ip", ip)

				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
			logger.Info(ctx, "Success")
		})
	}
}
