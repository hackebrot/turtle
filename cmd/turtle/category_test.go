package main

import "testing"

func TestRunCategory(t *testing.T) {
	tests := []cmdTestCase{{
		name:      "error missing category",
		args:      []string{"category"},
		wantError: true,
	}, {
		name:      "animals_and_nature",
		args:      []string{"category", "animals_and_nature"},
		outFile:   "category/animals_and_nature.json",
		wantError: false,
	}}
	runTestCmd(t, tests)
}
