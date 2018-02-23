package commands

import (
	"fmt"
	"io"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Delete implements the Delete command.
// It deletes a given key.
type Delete struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Delete) Do(stdout, stderr io.Writer, args []string) error {
	if args == nil || len(args) != 1 {
		_, err := stderr.Write([]byte(
			fmt.Sprintf("%s requires 1 arguments. Type HELP for more information.", cmd.Token()),
		))
		return err
	}

	if err := cmd.Stack.Peek().Store().Delete(args[0]); err != nil {
		_, err := stderr.Write([]byte(err.Error() + "\n"))
		return err
	}
	return nil
}

// Token implements Command
func (cmd *Delete) Token() string {
	return "DELETE"
}
