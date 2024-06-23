package commands

import (
	"fmt"
	"os"
	"strings"
)

func Chdir(args ...string) error {
	target := args[0]
	if strings.HasPrefix(target, "~") {
		target = strings.Replace(target, "~", os.Getenv("HOME"), 1)
	}
	err := os.Chdir(target)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", target)
		return err
	}

	return nil
}
