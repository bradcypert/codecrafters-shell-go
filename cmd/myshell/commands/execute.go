package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/util"
)

func Execute(target string, args ...string) error {
	pathItems, _ := util.PathSearch()
	for _, item := range pathItems {
		_, path := filepath.Split(item)
		if path == target {
			cmd := exec.Command(item, args...)
			out, err := cmd.Output()
			if err != nil {
				return err
			}
			fmt.Print(string(out))
			return nil
		}
	}
	return errors.New("executable not found")
}
