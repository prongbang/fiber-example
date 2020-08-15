package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/prongbang/fiber-example/internal/gofiber/todo"
	"github.com/prongbang/fiber-example/internal/pkg/ratelimit"
	"github.com/prongbang/fiber-example/pkg/middleware"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(middleware.SayHi())

	// Rate Limiter
	app.Use(ratelimit.Default())

	// Router
	todo.NewRouter(todo.NewHandler()).Register(app)

	app.Listen(3000)
}
