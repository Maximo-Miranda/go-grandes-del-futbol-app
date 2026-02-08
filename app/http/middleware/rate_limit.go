package middleware

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/goravel/framework/contracts/http"
)

type rateLimitEntry struct {
	count     int
	expiresAt time.Time
}

type rateLimiter struct {
	entries map[string]*rateLimitEntry
	mu      sync.RWMutex
}

var limiter = &rateLimiter{
	entries: make(map[string]*rateLimitEntry),
}

func init() {
	go limiter.cleanup()
}

func (r *rateLimiter) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		r.mu.Lock()
		now := time.Now()
		for key, entry := range r.entries {
			if now.After(entry.expiresAt) {
				delete(r.entries, key)
			}
		}
		r.mu.Unlock()
	}
}

func (r *rateLimiter) check(key string, maxAttempts int, decayMinutes int) (bool, int, time.Time) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	entry, exists := r.entries[key]

	if !exists || now.After(entry.expiresAt) {
		r.entries[key] = &rateLimitEntry{
			count:     1,
			expiresAt: now.Add(time.Duration(decayMinutes) * time.Minute),
		}
		return true, maxAttempts - 1, r.entries[key].expiresAt
	}

	if entry.count >= maxAttempts {
		return false, 0, entry.expiresAt
	}

	entry.count++
	return true, maxAttempts - entry.count, entry.expiresAt
}

func RateLimit(maxAttempts int, decayMinutes int) http.Middleware {
	return func(ctx http.Context) {
		key := fmt.Sprintf("rate_limit:%s:%s", ctx.Request().Ip(), ctx.Request().Path())
		allowed, remaining, retryAfter := limiter.check(key, maxAttempts, decayMinutes)

		ctx.Response().Header("X-RateLimit-Limit", strconv.Itoa(maxAttempts))
		ctx.Response().Header("X-RateLimit-Remaining", strconv.Itoa(remaining))

		if !allowed {
			retrySeconds := strconv.FormatInt(int64(time.Until(retryAfter).Seconds()), 10)
			ctx.Response().Header("Retry-After", retrySeconds)

			if ctx.Request().Header("X-Inertia") == "true" {
				ctx.Response().Json(http.StatusTooManyRequests, http.Json{
					"error":   "Too Many Requests",
					"message": "Demasiados intentos. Por favor espera un momento.",
				}).Abort()
			} else {
				ctx.Response().String(http.StatusTooManyRequests, "Demasiados intentos. Por favor espera un momento.").Abort()
			}
			return
		}

		ctx.Request().Next()
	}
}

func RateLimitAuth() http.Middleware {
	return RateLimit(30, 1)
}
