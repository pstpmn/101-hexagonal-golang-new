package main

import "github.com/gofiber/fiber/v2"

func InitRoute(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
