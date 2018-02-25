package commands

import (
	"github.com/frankgreco/fetakv/pkg/stack"
)

// Commit implements the Commit command.
// It commits a transaction.
type Commit struct {
	Stack *stack.Stack
}

// Do implements Command
func (cmd *Commit) Do(args []string) (stdout, stderr string) {
	if cmd.Stack.Size() > 1 {
		curr := cmd.Stack.Pop()
		cmd.Stack.Peek().Store().AddAll(curr.Store())
		return noOutput, noOutput
	}

	return noOutput, "There are no current transactions to commit."
}

// Token implements Command
func (cmd *Commit) Token() string {
	return "COMMIT"
}

// IsTerminal implements Command
func (cmd *Commit) IsTerminal() bool {
	return false
}
