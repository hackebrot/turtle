package main

import (
	"testing"
)

func TestRunCategory(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "error",
			args:    []string{"nope"},
			wantErr: true,
		},
		{
			name:    "no error",
			args:    []string{"people"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runCategory(nil, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("runCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
