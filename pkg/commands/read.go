package commands

import (
	"fmt"
	"io"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Read implements the Read command.
// It reads the value associated with the given key.
type Read struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Read) Do(stdout, stderr io.Writer, args []string) error {
	if args == nil || len(args) != 1 {
		_, err := stderr.Write([]byte(
			fmt.Sprintf("%s requires 1 arguments. Type HELP for more information.", cmd.Token()),
		))
		return err
	}

	val, err := cmd.Stack.Peek().Store().Read(args[0])
	if err != nil {
		_, err := stderr.Write([]byte(err.Error() + "\n"))
		return err
	}

	if _, err := stdout.Write([]byte(val + "\n")); err != nil {
		return err
	}

	return nil
}

// Token implements Command
func (cmd *Read) Token() string {
	return "READ"
}
