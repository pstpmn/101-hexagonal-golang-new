package pkg

import (
	"testing"
	"time"
)

func Test_j_Generate(t *testing.T) {
	type args struct {
		data map[string]interface{}
		key  string
		exp  time.Timer
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := j{}
			got, err := j.Generate(tt.args.data, tt.args.key, tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
