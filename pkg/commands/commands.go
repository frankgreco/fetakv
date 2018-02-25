package commands

import (
	"errors"
	"fmt"
	"strings"
)

// Commands represents the list of valid commands.
type Commands []Command

// Command provides an interface for a command
type Command interface {
	// Token returns token corresponding to this command.
	Token() string
	// Do preforms any logic associated with this command.
	// If this logic results in information that should be
	// relayed to stdout or stderr, it is returned.
	Do(args []string) (stdout, stderr string)
	// IsTerminal relays whether this program should exit
	// after the exectution of Do.
	IsTerminal() bool
}

var noOutput string // note that zero value is "" which is what we would've set it too

// Register will register a list of commands as valid.
func Register(commands ...Command) *Commands {
	cmds := new(Commands)
	for _, item := range commands {
		*cmds = append(*cmds, item)
	}
	return cmds
}

// Parse will attempt to find a registered command that
// can be identified by the given token.
func (cmds *Commands) Parse(data string) (Command, error) {
	for _, cmd := range *cmds {
		if cmd.Token() == strings.ToUpper(data) {
			return cmd, nil
		}
	}
	msg := fmt.Sprintf("Command not found: %s", data) // bypass golint
	return nil, errors.New(msg)
}
