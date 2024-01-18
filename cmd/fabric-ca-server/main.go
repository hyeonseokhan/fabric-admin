package main

import (
	"os"
)

var (
	blockingStart = true
)

func main() {
	if err := RunManin(os.Args); err != nil {
		os.Exit(1)
	}
}

func RunManin(args []string) error {
	saveOsArgs := os.Args
	os.Args = args

	cmdName := ""
	if len(args) > 1 {
		cmdName = args[1]
	}
	scmd := NewCommand(cmdName, blockingStart)

	err := scmd.Execute()

	os.Args = saveOsArgs

	return nil
}
