package commands

import (
	"github.com/frankgreco/fetakv/pkg/stack"
)

// Abort implements the Abort command.
// It aborts a transaction.
type Abort struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Abort) Do(args []string) (stdout, stderr string) {
	if cmd.Stack.Size() > 1 {
		cmd.Stack.Pop()
		return noOutput, noOutput
	}

	return noOutput, "There are no current transactions to abort."
}

// Token implements Command
func (cmd *Abort) Token() string {
	return "ABORT"
}

// IsTerminal implements Command
func (cmd *Abort) IsTerminal() bool {
	return false
}
