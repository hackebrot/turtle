package main

import (
	"testing"
)

func TestRunKeyword(t *testing.T) {
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
			args:    []string{"dog"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runKeyword(nil, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("runKeyword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
