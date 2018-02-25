package commands

import (
	"github.com/frankgreco/fetakv/pkg/stack"
	"github.com/frankgreco/fetakv/pkg/transaction"
)

// Start implements the Start command.
// It starts a transation.
type Start struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Start) Do(args []string) (stdout, stderr string) {
	cmd.Stack.Push(transaction.New())
	return noOutput, noOutput
}

// Token implements Command
func (cmd *Start) Token() string {
	return "START"
}

// IsTerminal implements Command
func (cmd *Start) IsTerminal() bool {
	return false
}
