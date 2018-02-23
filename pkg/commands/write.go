package commands

import (
	"fmt"
	"io"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Write implements the Write command.
// It store a value in a key.
type Write struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Write) Do(stdout, stderr io.Writer, args []string) error {
	if args == nil || len(args) != 2 {
		_, err := stderr.Write([]byte(
			fmt.Sprintf("%s requires 2 arguments. Type HELP for more information.", cmd.Token()),
		))
		return err
	}

	cmd.Stack.Peek().Store().Write(args[0], args[1])
	return nil
}

// Token implements Command
func (cmd *Write) Token() string {
	return "WRITE"
}
