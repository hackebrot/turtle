package main

import "testing"

func TestRunRandom(t *testing.T) {
	tests := []cmdTestCase{{
		name:      "random",
		args:      []string{"random"},
		wantError: false,
	}}
	runTestCmd(t, tests)
}
