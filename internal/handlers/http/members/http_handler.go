package handlers

import (
	"fmt"
	"lean-oauth/internal/core/ports"
	"time"

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

	dob, error := time.Parse("2006-01-02", req.Dob)
	if error != nil {
		return hdl.response.Json(c, fiber.StatusBadRequest, fmt.Sprint("invalid birth day"), nil, false)
	}

	if err := req.Validate(); err != nil {
		return hdl.response.Json(c, fiber.StatusBadRequest, fmt.Sprint(err), nil, false)
	}

	member, err := hdl.membersUseCase.NewMember(req.Username, req.Password, req.FistName, req.LastName, dob)
	if err != nil {
		return hdl.response.Json(c, fiber.StatusOK, fmt.Sprint(err), nil, false)
	}

	result := map[string]interface{}{
		"username":  member.Username,
		"createdAt": member.CreatedAt,
	}

	return hdl.response.Json(c, fiber.StatusCreated, "register successful", result, true)
}
