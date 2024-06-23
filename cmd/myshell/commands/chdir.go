package commands

import (
	"fmt"
	"os"
)

func Chdir(args ...string) error {
	target := args[0]
	err := os.Chdir(target)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", target)
		return err
	}

	return nil
}
