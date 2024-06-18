package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Wait for user input
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := stdin.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
	}

}
