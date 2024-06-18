package commands

import (
	"fmt"
	"strings"
)

func Echo(args ...string) {
	fmt.Printf("%s\n", strings.Join(args, " "))
}
