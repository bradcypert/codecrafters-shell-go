package commands

import (
	"fmt"
	"os"
)

func Chdir(args ...string) error {
	target := args[0]
	err := os.Chdir(target)
	if err != nil {
		fmt.Println("cd: %s: No such file or directory", target)
		return err
	}

	return nil
}
