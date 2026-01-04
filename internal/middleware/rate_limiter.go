package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiterConfig holds configuration for rate limiting
type RateLimiterConfig struct {
	Max        int           // Maximum number of requests
	Expiration time.Duration // Time window for requests
}

// Default rate limiting: 100 requests per 15 minutes
var DefaultConfig = RateLimiterConfig{
	Max:        100,
	Expiration: 15 * time.Minute,
}

// RateLimiter returns a rate limiting middleware
func RateLimiter(config RateLimiterConfig) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        config.Max,
		Expiration: config.Expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			// Use IP address as the key
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later",
			})
		},
	})
}

// AuthRateLimiter stricter rate limiting for auth endpoints
func AuthRateLimiter() fiber.Handler {
	return RateLimiter(RateLimiterConfig{
		Max:        5,           // 5 requests
		Expiration: time.Minute, // per minute
	})
}
