package commands

import (
	"io"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Abort implements the Abort command.
// It aborts a transaction.
type Abort struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Abort) Do(stdout, stderr io.Writer, args []string) error {
	if cmd.Stack.Size() > 1 {
		cmd.Stack.Pop()
		return nil
	}
	_, err := stderr.Write([]byte("There are no current transactions to abort.\n"))
	return err
}

// Token implements Command
func (cmd *Abort) Token() string {
	return "ABORT"
}
