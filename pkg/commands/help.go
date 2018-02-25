package commands

// Help implements the Help command.
// It displays usage information.
type Help struct{}

// Do implements Command
func (cmd *Help) Do(args []string) (stdout, stderr string) {
	return `fetakv a command line REPL that drives an in-memory kv storage system.

Available Commands:
  HELP                Display usage information.
  READ    <key>       Read the value associated with the given key.
  WRITE   <key> <val> Stores val in key.
  DELETE  <key>       Removes key from the store.
  START               Start a transation (nested transactions supported).
  COMMIT              Commit a transaction.
  ABORT               Aborts a transation.
  QUIT                Exit fetakv.`, noOutput
}

// Token implements Command
func (cmd *Help) Token() string {
	return "HELP"
}

// IsTerminal implements Command
func (cmd *Help) IsTerminal() bool {
	return false
}
