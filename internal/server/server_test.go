package server

import (
	"lean-oauth/internal/core/ports"
	"testing"
)

func Test_server_Initialize(t *testing.T) {
	type fields struct {
		membersHandler ports.IMembersHandler
		env            map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := server{
				membersHandler: tt.fields.membersHandler,
				env:            tt.fields.env,
			}
			s.Initialize()
		})
	}
}
