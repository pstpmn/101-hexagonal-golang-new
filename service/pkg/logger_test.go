//go:build integration
// +build integration

package pkg

import (
	"testing"
)

func Test_logger_Log(t *testing.T) {
	// new object
	logger := NewLogger()
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test integrate logging",
			args{msg: "say hi my friend"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger.Log(tt.args.msg)
		})
	}
}

func Test_logger_error(t *testing.T) {
	// new object
	logger := NewLogger()
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test integrate logging error",
			args{msg: "say hi my error code"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger.Error(tt.args.msg)
		})
	}
}
