package main

import "testing"

func TestRunSearch(t *testing.T) {
	tests := []cmdTestCase{{
		name:      "error",
		args:      []string{"search", "aaaaaaaaaa"},
		wantError: true,
	}, {
		name:      "single match",
		args:      []string{"search", "rocket"},
		outFile:   "search/rocket.json",
		wantError: false,
	}, {
		name:      "multiple matches",
		args:      []string{"search", "dog"},
		outFile:   "search/dog.json",
		wantError: false,
	}}
	runTestCmd(t, tests)
}
