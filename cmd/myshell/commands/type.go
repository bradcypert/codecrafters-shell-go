package commands

import (
	"fmt"
)

const builtin = " is a shell builtin"
const notFound = ": not found"

func Type(args ...string) {
	builtins := [...]string{"echo", "exit", "type"}
	target := args[0]

	for _, b := range builtins {
		if b == target {
			fmt.Printf("%s%s", target, builtin)
			return
		}
	}

	fmt.Printf("%s%s", target, notFound)
}
