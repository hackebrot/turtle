package main

import "testing"

func TestRunChar(t *testing.T) {
	tests := []cmdTestCase{{
		name:      "error missing char",
		args:      []string{"char"},
		wantError: true,
	}, {
		name:      "error unknown char",
		args:      []string{"char", "abc"},
		wantError: true,
	}, {
		name:      "robot",
		args:      []string{"char", "🤖"},
		outFile:   "char/robot.json",
		wantError: false,
	}}
	runTestCmd(t, tests)
}
