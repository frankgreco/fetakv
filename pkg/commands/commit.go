package commands

import (
	"io"

	"github.com/frankgreco/fetakv/pkg/stack"
)

// Commit implements the Commit command.
// It commits a transaction.
type Commit struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Commit) Do(stdout, stderr io.Writer, args []string) error {
	if cmd.Stack.Size() > 1 {
		curr := cmd.Stack.Pop()
		cmd.Stack.Peek().Store().AddAll(curr.Store())
		return nil
	}

	_, err := stderr.Write([]byte("There are no current transactions to commit.\n"))
	return err
}

// Token implements Command
func (cmd *Commit) Token() string {
	return "COMMIT"
}
