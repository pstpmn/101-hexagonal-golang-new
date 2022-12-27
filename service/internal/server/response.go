package server

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"msg,omitempty"`
	Status  bool        `json:"status"`
	Result  interface{} `json:"result,omitempty"`
}

type AuthorizeResponse struct {
	Message    string      `json:"msg,omitempty"`
	Status     bool        `json:"status"`
	Authorized bool        `json:"authorized"`
	Result     interface{} `json:"result,omitempty"`
}

type IResponse interface {
	Json(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool) error
	ErrorRequestBody(h *fiber.Ctx) error
	JsonAuth(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool, isValidAuthorize bool) error
}

type y struct {
}

func NewResponse() IResponse {
	return &y{}
}

func (y y) JsonAuth(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool, isValidAuthorize bool) error {
	h.Set("Content-Type", "application/json")
	res := AuthorizeResponse{
		Message:    message,
		Result:     result,
		Status:     status,
		Authorized: isValidAuthorize,
	}
	return h.Status(httpCode).JSON(res)
}

func (y y) Json(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool) error {
	h.Set("Content-Type", "application/json")
	res := Response{
		Message: message,
		Result:  result,
		Status:  status,
	}
	return h.Status(httpCode).JSON(res)
}

func (y y) ErrorRequestBody(h *fiber.Ctx) error {
	h.Set("Content-Type", "application/json")
	res := Response{
		Message: "parse http body invalid",
		Status:  false,
	}
	return h.Status(fiber.ErrBadRequest.Code).JSON(res)
}

func (y y) Unauthorized(h *fiber.Ctx) error {
	h.Set("Content-Type", "application/json")
	res := Response{
		Message: "invalid token",
		Status:  false,
	}
	return h.Status(fiber.ErrBadRequest.Code).JSON(res)
}
