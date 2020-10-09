package main

import (
	"os"

	"github.com/lukitsbrian/poe-sg/cmd/poe-sgd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
