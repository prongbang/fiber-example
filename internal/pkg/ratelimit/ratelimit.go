package ratelimit

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
)

// 3 requests per 10 seconds max
func Default() func(*fiber.Ctx) {
	return limiter.New(limiter.Config{
		Timeout: 10,
		Max:     3,
	})
}
