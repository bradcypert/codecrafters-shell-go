package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const builtin = " is a shell builtin"
const notFound = ": not found"

func Type(args ...string) {
	builtins := [...]string{"echo", "exit", "type"}

	target := args[0]
	for _, b := range builtins {
		if b == target {
			fmt.Printf("%s%s\n", target, builtin)
			return
		}
	}

	// TODO: Probably refactor this out since we're going to need this
	// for executing programs too
	path := os.Getenv("PATH")

	if path != "" {
		// crawl path for more things to search
		pathFragments := strings.Split(path, ":")
		for _, f := range pathFragments {
			searchPath := filepath.Join(f, target)
			if _, err := os.Stat(searchPath); err == nil {
				fmt.Printf("%s is %s\n", target, searchPath)
				return
			}
		}
	}

	fmt.Printf("%s%s\n", target, notFound)
}
