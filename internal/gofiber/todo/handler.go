package todo

import (
	"net/http"

	"github.com/gofiber/fiber"
)

type Handler struct {
}

func (h Handler) GetById(c *fiber.Ctx) {
	c.JSON(&fiber.Map{
		"id": c.Params("id"),
	})
}

func (h Handler) PostTodo(c *fiber.Ctx) {
	todo := Todo{}
	if err := c.BodyParser(&todo); err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return
	}
	c.JSON(&todo)
}

func (h Handler) PutTodo(c *fiber.Ctx) {
	todo := Todo{}
	if err := c.BodyParser(&todo); err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return
	}
	c.JSON(&fiber.Map{
		"id":   c.Params("id"),
		"body": todo,
	})
}

func (h Handler) PatchTodo(c *fiber.Ctx) {
	todo := Todo{}
	if err := c.BodyParser(&todo); err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return
	}
	c.JSON(&fiber.Map{
		"id":   c.Params("id"),
		"body": todo,
	})
}

func (h Handler) DeleteTodo(c *fiber.Ctx) {
	c.JSON(&fiber.Map{
		"id": c.Params("id"),
	})
}

func NewHandler() Handler {
	return Handler{}
}
