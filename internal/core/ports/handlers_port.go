package ports

import "github.com/gofiber/fiber/v2"

type IMembersHandler interface {
	HelloWorld(c *fiber.Ctx) error
	Registration(c *fiber.Ctx) error
	Authentication(c *fiber.Ctx) error
}
