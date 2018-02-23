package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/frankgreco/fetakv/cmd/fetakv/app"
	"github.com/frankgreco/fetakv/pkg/utils"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := app.Run(os.Stdin, os.Stdout, os.Stderr, utils.PromptCaret); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
