package todo

import "github.com/gofiber/fiber"

type Router struct {
	Handle Handler
}

func (r Router) Register(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/todo/:id", r.Handle.GetById)
	v1.Post("/todo", r.Handle.PostTodo)
	v1.Put("/todo/:id", r.Handle.PutTodo)
	v1.Patch("/todo/:id", r.Handle.PatchTodo)
	v1.Delete("/todo/:id", r.Handle.DeleteTodo)
}

func NewRouter(handle Handler) Router {
	return Router{
		Handle: handle,
	}
}
