package commands

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Commands represents the list of valid commands.
type Commands []Command

// Command provides an interface for a command
type Command interface {
	Token() string
	Do(stdout, stderr io.Writer, args []string) error
}

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
	msg := fmt.Sprintf("Command not found: %s\n", data) // bypass golint
	return nil, errors.New(msg)
}
