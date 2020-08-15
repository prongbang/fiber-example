package middleware

import (
	"log"

	"github.com/gofiber/fiber"
)

func SayHi() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		log.Println("-> Hi")
		c.Next()
	}
}
