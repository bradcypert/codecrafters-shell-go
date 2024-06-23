package commands

import (
	"fmt"
	"strings"
)

func Echo(args ...string) error {
	fmt.Printf("%s\n", strings.Join(args, " "))
	return nil
}
