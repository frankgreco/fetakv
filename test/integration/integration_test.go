package integration

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/frankgreco/fetakv/cmd/fetakv/app"
	"github.com/frankgreco/fetakv/pkg/utils"
)

func TestIntegration(t *testing.T) {

	tests := []struct {
		name                                  string
		stdin, expectedStdout, expectedStderr []byte
	}{
		{
			name:           "simple",
			stdin:          simple.stdin,
			expectedStdout: simple.stdout,
			expectedStderr: simple.stderr,
		},
		{
			name:           "complex",
			stdin:          complex.stdin,
			expectedStdout: complex.stdout,
			expectedStderr: complex.stderr,
		},
	}

	for _, test := range tests {
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		stdoutWriter := bufio.NewWriter(stdout)
		stderrWriter := bufio.NewWriter(stderr)
		err := app.Run(bytes.NewReader(test.stdin), stdoutWriter, stderrWriter, utils.PromptNone)
		stdoutWriter.Flush()
		stderrWriter.Flush()

		if err != nil {
			fmt.Printf("expected nil received %v\n", err)
			t.Fail()
		}
		if !bytes.Equal(stdout.Bytes(), test.expectedStdout) {
			fmt.Printf("stdout expected %v received %v\n", test.expectedStdout, stdout.Bytes())
			t.Fail()
		}
		if !bytes.Equal(stderr.Bytes(), test.expectedStderr) {
			fmt.Printf("stderr expected %v received %v\n", test.expectedStderr, stderr.Bytes())
			t.Fail()
		}
	}
}
