package commands

// Quit implements the Quit command.
// It quits the program.
type Quit struct{}

// Do implements Command
func (cmd *Quit) Do(args []string) (stdout, stderr string) {
	return "Exiting...", noOutput
}

// Token implements Command
func (cmd *Quit) Token() string {
	return "QUIT"
}

// IsTerminal implements Command
func (cmd *Quit) IsTerminal() bool {
	return true
}
