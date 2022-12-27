package ports

import "github.com/gofiber/fiber/v2"

type IMiddlewares interface {
	Authorize(h *fiber.Ctx) error
}
