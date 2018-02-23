package commands

import (
	"io"
	"os"
)

// Quit implements the Quit command.
// It quits the program.
type Quit struct{}

// Do implements Command
func (cmd *Quit) Do(stdout, stderr io.Writer, args []string) error {
	if _, err := stdout.Write([]byte("Exiting...\n")); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
	return nil
}

// Token implements Command
func (cmd *Quit) Token() string {
	return "QUIT"
}
