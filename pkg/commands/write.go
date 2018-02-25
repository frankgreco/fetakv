package commands

import (
	"fmt"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Write implements the Write command.
// It store a value in a key.
type Write struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Write) Do(args []string) (stdout, stderr string) {
	if args == nil || len(args) != 2 {
		return noOutput, fmt.Sprintf("%s requires 2 arguments. Type HELP for more information.", cmd.Token())
	}

	cmd.Stack.Peek().Store().Write(args[0], args[1])
	return noOutput, noOutput
}

// Token implements Command
func (cmd *Write) Token() string {
	return "WRITE"
}

// IsTerminal implements Command
func (cmd *Write) IsTerminal() bool {
	return false
}
