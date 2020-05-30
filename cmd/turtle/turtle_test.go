package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hackebrot/go-repr/repr"
)

// cmdTestCase describes a test case for turtle commands
type cmdTestCase struct {
	name      string
	args      []string
	outFile   string
	wantError bool
}

const testdata = "testdata"

// runTestCmd is a test helper that runs the given test cases
func runTestCmd(t *testing.T, tests []cmdTestCase) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("💻 running 'turtle %s'", strings.Join(tt.args, " "))

			out, err := executeCmd(tt.args)

			if tt.wantError && err == nil {
				t.Fatalf("🤖 command did not return error")
			}

			if !tt.wantError && err != nil {
				t.Errorf("🤖 unexpected error: %v", err)
			}

			if tt.outFile != "" {
				data, err := ioutil.ReadFile(filepath.Join(testdata, tt.outFile))
				if err != nil {
					t.Fatalf("🤖 error reading outFile: %v", err)
				}

				got := strings.TrimSpace(out)
				want := strings.TrimSpace(string(data))

				if !cmp.Equal(got, want) {
					t.Errorf(
						"📋 cmd output does not match '%s'\n\ngot:\n%s\n\nwant:\n%s\n",
						tt.outFile,
						repr.Repr(got),
						repr.Repr(want),
					)
				}
			}
		})
	}
}

// executeCmd creates and executes a new turtle cmd
func executeCmd(args []string) (string, error) {
	buf := new(bytes.Buffer)
	root := newTurtleCmd(buf)
	root.SetOut(buf)
	root.SetArgs(args)

	err := root.Execute()

	return buf.String(), err
}

func TestVersionFlag(t *testing.T) {
	tests := []cmdTestCase{{
		name:      "version",
		args:      []string{"--version"},
		wantError: false,
	}}
	runTestCmd(t, tests)
}
