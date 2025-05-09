package todo

import "github.com/gofiber/fiber/v2"

type TodoHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	t := &TodoHandler{
		router: router,
	}

	api := t.router.Group("/todo")
	api.Get("/", t.getAll)
}

func (t *TodoHandler) getAll(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
