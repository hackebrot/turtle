package main

import (
	"testing"
)

func Test_runChar(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "error",
			args:    []string{"nopeasdasdasdsd"},
			wantErr: true,
		},
		{
			name:    "no error",
			args:    []string{"🤖"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runChar(nil, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("runChar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
