package commands

import (
	"os"
)

func Chdir(args ...string) error {
	target := args[0]
	err := os.Chdir(target)
	if err != nil {
		return err
	}

	return nil
}
