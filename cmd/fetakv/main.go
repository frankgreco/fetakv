package main

import (
	"flag"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/frankgreco/fetakv/cmd/fetakv/app"
	"github.com/frankgreco/fetakv/pkg/prompt"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	stdin := flag.String("stdin", "", "stdin")
	flag.Parse()

	r, p, err := getStdin(*stdin)
	if err != nil {
		os.Exit(1)
	}

	if err := app.Run(r, os.Stdout, os.Stderr, p); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func getStdin(stdin string) (io.Reader, prompt.Prompt, error) {
	if len(stdin) < 1 {
		return os.Stdin, prompt.PromptCaret, nil
	}
	r, err := os.Open(stdin)
	if err != nil {
		return nil, "", err
	}
	return r, prompt.PromptNone, nil
}
