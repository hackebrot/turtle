package main

import "testing"

func TestRunKeyword(t *testing.T) {
	tests := []cmdTestCase{{
		name:      "error missing keyword",
		args:      []string{"keyword"},
		wantError: true,
	}, {
		name:      "error unknown keyword",
		args:      []string{"keyword", "foo"},
		wantError: true,
	}, {
		name:      "mountain",
		args:      []string{"keyword", "mountain"},
		outFile:   "keyword/mountain.json",
		wantError: false,
	}}
	runTestCmd(t, tests)
}
