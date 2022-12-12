package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lean-oauth/internal/core/ports"
)

type HTTPHandler struct {
	membersUseCase ports.MembersUseCase
}

func NewHTTPHandler(membersUseCase ports.MembersUseCase) *HTTPHandler {
	return &HTTPHandler{membersUseCase: membersUseCase}
}

func (hdl *HTTPHandler) HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
