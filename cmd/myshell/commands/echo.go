package commands

import (
	"fmt"
	"strings"
)

func Echo(args ...string) {
	fmt.Print(strings.Join(args, " "))
}
