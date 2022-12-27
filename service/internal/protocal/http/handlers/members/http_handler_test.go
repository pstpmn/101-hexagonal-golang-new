package handlers

import (
	"github.com/gofiber/fiber/v2"
	ports2 "learn-oauth2/internal/core/ports"
	"testing"
)

func TestHTTPHandler_Registration(t *testing.T) {
	type fields struct {
		membersUseCase ports2.MembersUseCase
		response       ports2.IResponse
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

func TestHTTPHandler_HelloWorld(t *testing.T) {
	type fields struct {
		membersUseCase    ports2.MembersUseCase
		response          ports2.IResponse
		authenticationKey string
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
				membersUseCase:    tt.fields.membersUseCase,
				response:          tt.fields.response,
				authenticationKey: tt.fields.authenticationKey,
			}
			if err := hdl.HelloWorld(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("HelloWorld() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
