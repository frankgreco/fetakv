package commands

import (
	"fmt"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Read implements the Read command.
// It reads the value associated with the given key.
type Read struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Read) Do(args []string) (stdout, stderr string) {
	if args == nil || len(args) != 1 {
		return noOutput, fmt.Sprintf("%s requires 1 arguments. Type HELP for more information.", cmd.Token())
	}

	val, err := cmd.Stack.Peek().Store().Read(args[0])
	if err != nil {
		return noOutput, err.Error()
	}

	return val, noOutput
}

// Token implements Command
func (cmd *Read) Token() string {
	return "READ"
}

// IsTerminal implements Command
func (cmd *Read) IsTerminal() bool {
	return false
}
