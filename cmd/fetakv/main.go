package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/frankgreco/fetakv/cmd/fetakv/app"
	"github.com/frankgreco/fetakv/pkg/prompt"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := app.Run(os.Stdin, os.Stdout, os.Stderr, prompt.PromptCaret); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
