package commands

import (
	"fmt"
	"os"
)

func Pwd(args ...string) error {
	path, err := os.Executable()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", path)

	return nil
}
