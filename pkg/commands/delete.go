package commands

import (
	"fmt"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Delete implements the Delete command.
// It deletes a given key.
type Delete struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Delete) Do(args []string) (stdout, stderr string) {
	if args == nil || len(args) != 1 {
		return noOutput, fmt.Sprintf("%s requires 1 arguments. Type HELP for more information.", cmd.Token())
	}

	if err := cmd.Stack.Peek().Store().Delete(args[0]); err != nil {
		return noOutput, err.Error()
	}
	return noOutput, noOutput
}

// Token implements Command
func (cmd *Delete) Token() string {
	return "DELETE"
}

// IsTerminal implements Command
func (cmd *Delete) IsTerminal() bool {
	return false
}
