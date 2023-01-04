package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	ports2 "learn-oauth2/internal/core/ports"
	"time"
)

type HTTPHandler struct {
	membersUseCase    ports2.MembersUseCase
	response          ports2.IResponse
	authenticationKey string
}

func NewHTTPHandler(membersUseCase ports2.MembersUseCase, response ports2.IResponse, key string) ports2.IMembersHandler {
	return &HTTPHandler{membersUseCase: membersUseCase, response: response, authenticationKey: key}
}

func (hdl *HTTPHandler) Authorization(c *fiber.Ctx) error {
	clientToken := c.Get("Authorize-token")
	if clientToken == "" {
		return hdl.response.Json(c, fiber.StatusBadRequest, "empty authorize token", nil, false)
	}

	if _, err := hdl.membersUseCase.Authorization(clientToken, hdl.authenticationKey); err != nil {
		return hdl.response.Json(c, fiber.StatusBadRequest, fmt.Sprintf("%s", err), nil, false)
	}

	return hdl.response.Json(c, fiber.StatusOK, "authorize successful", nil, true)
}

func (hdl *HTTPHandler) HelloWorld(c *fiber.Ctx) error {
	result := map[string]interface{}{
		"username":    c.Locals("username"),
		"firstName":   c.Locals("firstName"),
		"lastName":    c.Locals("lastName"),
		"dateOfBirth": c.Locals("dob"),
		"createdAt":   c.Locals("createdAt"),
	}
	return hdl.response.JsonAuth(c, fiber.StatusOK, fmt.Sprintf("welcome you : %s %s", c.Locals("firstName"), c.Locals("lastName")), result, true, true)
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

func (hdl *HTTPHandler) Authentication(c *fiber.Ctx) error {
	req := &AuthenticationRequest{}
	if err := c.BodyParser(req); err != nil {
		return hdl.response.ErrorRequestBody(c)
	}

	token, mem, err := hdl.membersUseCase.Authentication(req.Username, req.Password, hdl.authenticationKey)
	if err != nil {
		return hdl.response.Json(c, fiber.StatusOK, fmt.Sprint(err), nil, false)
	}

	result := map[string]interface{}{
		"token":     token,
		"username":  mem.Username,
		"firstName": mem.FirstName,
		"lastName":  mem.LastName,
	}
	return hdl.response.Json(c, fiber.StatusOK, "authentication successful", result, true)
}
