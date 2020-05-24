package main

import (
	"testing"
)

func TestRunRandom(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "success",
			args:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runRandom(nil, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("runRandom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
