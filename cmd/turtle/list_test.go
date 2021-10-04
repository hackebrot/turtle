package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hackebrot/go-repr/repr"
)

func TestAddIfUnique(t *testing.T) {
	type args struct {
		r      []string
		sItems []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "one element",
			args: args{
				r:      []string{},
				sItems: []string{"hello"},
			},
			want: []string{"hello"},
		},
		{
			name: "non empty",
			args: args{
				r:      []string{"hello"},
				sItems: []string{"world", "zzz"},
			},
			want: []string{"hello", "world", "zzz"},
		},
		{
			name: "duplicate",
			args: args{
				r:      []string{"hello"},
				sItems: []string{"world", "hello"},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "sorted",
			args: args{
				r:      []string{"hello"},
				sItems: []string{"world", "a"},
			},
			want: []string{"a", "hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addIfUnique(tt.args.r, tt.args.sItems...); !cmp.Equal(got, tt.want) {
				t.Errorf("addIfUnique() = %v, want %v", repr.Repr(got), repr.Repr(tt.want))
			}
		})
	}
}
func TestRunList(t *testing.T) {
	tests := []cmdTestCase{{
		name:      "show help for missing subcommand",
		args:      []string{"list"},
		wantError: false,
	}, {
		name:      "categories",
		args:      []string{"list", "categories"},
		outFile:   "list/categories.json",
		wantError: false,
	}, {
		name:      "keywords",
		args:      []string{"list", "keywords"},
		wantError: false,
	}, {
		name:      "names",
		args:      []string{"list", "names"},
		wantError: false,
	}}
	runTestCmd(t, tests)
}
