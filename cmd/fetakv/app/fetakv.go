package app

import (
	"bufio"
	"io"
	"strings"

	"github.com/frankgreco/fetakv/pkg/commands"
	"github.com/frankgreco/fetakv/pkg/stack"
	"github.com/frankgreco/fetakv/pkg/transaction"
	"github.com/frankgreco/fetakv/pkg/utils"
)

// Run fascilites this program
func Run(stdin io.Reader, stdout, stderr io.Writer, prompt utils.Prompt) error {
	s := stack.New()
	s.Push(transaction.New())

	registeredCommands := commands.Register(
		&commands.Read{Stack: s},
		&commands.Write{Stack: s},
		&commands.Delete{Stack: s},
		&commands.Start{Stack: s},
		&commands.Commit{Stack: s},
		&commands.Abort{Stack: s},
		&commands.Quit{},
		&commands.Help{},
	)

	scanner := bufio.NewScanner(stdin)
	scanner.Split(bufio.ScanLines)

	for {
		if _, err := stdout.Write([]byte(prompt)); err != nil {
			return err
		}
		if !scanner.Scan() {
			break
		}
		args := strings.Fields(scanner.Text())
		if len(args) < 1 {
			continue
		}

		cmd, err := registeredCommands.Parse(args[0])
		if err != nil {
			if _, err := stderr.Write([]byte(err.Error())); err != nil {
				return err
			}
			continue
		}
		if err := cmd.Do(stdout, stderr, args[1:]); err != nil {
			return err
		}
	}

	return scanner.Err()
}
