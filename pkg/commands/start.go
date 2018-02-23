package commands

import (
	"io"

	"github.com/frankgreco/fetakv/pkg/stack"
	"github.com/frankgreco/fetakv/pkg/transaction"
)

// Start implements the Start command.
// It starts a transation.
type Start struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Start) Do(stdout, stderr io.Writer, args []string) error {
	cmd.Stack.Push(transaction.New())
	return nil
}

// Token implements Command
func (cmd *Start) Token() string {
	return "START"
}
