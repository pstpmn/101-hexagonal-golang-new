package handlers

import (
	"fmt"
	"lean-oauth/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	membersUseCase ports.MembersUseCase
	response       ports.IResponse
}

func NewHTTPHandler(membersUseCase ports.MembersUseCase, response ports.IResponse) ports.IMembersHandler {
	return &HTTPHandler{membersUseCase: membersUseCase, response: response}
}

func (hdl *HTTPHandler) HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (hdl *HTTPHandler) Registration(c *fiber.Ctx) error {
	req := &RegistrationRequest{}
	if err := c.BodyParser(req); err != nil {
		return hdl.response.ErrorRequestBody(c)
	}

	if err := req.Validate(); err != nil {
		return hdl.response.Json(c, fiber.StatusBadRequest, fmt.Sprint(err), nil, false)
	}

	return hdl.response.ErrorRequestBody(c)
}
