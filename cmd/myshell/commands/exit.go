package commands

import (
	"os"
	"strconv"
)

func Exit(args ...string) error {
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		os.Exit(0)
	}
	os.Exit(exitCode)

	return nil
}
