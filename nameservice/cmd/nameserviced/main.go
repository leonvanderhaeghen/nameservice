package main

import (
	"os"

	"github.com/leonvanderhaeghen/nameservice/cmd/nameserviced/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
