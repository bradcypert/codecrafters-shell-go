package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
)

func main() {

	// Wait for user input
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		c, err := stdin.ReadString('\n')
		c = strings.TrimRight(c, "\n")
		if err != nil {
			os.Exit(1)
		}

		cmd := strings.Split(c, " ")
		command := cmd[0]
		args := cmd[1:]

		switch command {
		case "exit":
			commands.Exit(args...)
		case "echo":
			commands.Echo(args...)
		case "type":
			commands.Type(args...)
		case "pwd":
			commands.Pwd(args...)

		default:
			// assume trying to execute external program
			err := commands.Execute(command, args...)
			if err != nil {
				fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
			}
		}
	}

}
