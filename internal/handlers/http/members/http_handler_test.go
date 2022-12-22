package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lean-oauth/internal/core/ports"
	"testing"
)

func TestHTTPHandler_Registration(t *testing.T) {
	type fields struct {
		membersUseCase ports.MembersUseCase
		response       ports.IResponse
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hdl := &HTTPHandler{
				membersUseCase: tt.fields.membersUseCase,
				response:       tt.fields.response,
			}
			if err := hdl.Registration(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Registration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
