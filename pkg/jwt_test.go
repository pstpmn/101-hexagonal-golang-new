//go:build integration
// +build integration

package pkg

import (
	"testing"
	"time"
)

func Test_j_Generate(t *testing.T) {
	type args struct {
		data map[string]interface{}
		key  string
		exp  time.Time
	}
	tests := []struct {
		name string
		args args
		//want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test created token should be success",
			args: args{exp: time.Now().Add(time.Hour * 3), key: "secret", data: map[string]interface{}{"test": "2"}},
			//want:    "token",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := j{}
			_, err := j.Generate(tt.args.data, tt.args.key, tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
