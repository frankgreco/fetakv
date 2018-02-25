package app

import (
	"bufio"
	"io"
	"strings"

	"github.com/frankgreco/fetakv/pkg/commands"
	"github.com/frankgreco/fetakv/pkg/prompt"
	"github.com/frankgreco/fetakv/pkg/stack"
	"github.com/frankgreco/fetakv/pkg/transaction"
)

// Run fascilites this program
func Run(stdin io.Reader, stdout, stderr io.Writer, p prompt.Prompt) error {
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
		if err := write(p.String(), stdout); err != nil {
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
			if err := write(normalize(err.Error()), stderr); err != nil {
				return err
			}
			continue
		}
		stdoutText, stderrText := cmd.Do(args[1:])
		if err := write(normalize(stdoutText), stdout); err != nil {
			return err
		}
		if err := write(normalize(stderrText), stderr); err != nil {
			return err
		}
		if cmd.IsTerminal() {
			return nil
		}
	}

	return scanner.Err()
}

func write(data string, writer io.Writer) error {
	_, err := writer.Write([]byte(data))
	return err
}

func normalize(data string) string {
	if len(data) < 1 {
		return data
	}
	return data + "\n"
}
