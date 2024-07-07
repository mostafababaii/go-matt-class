package hello

import (
	"fmt"
	"strings"
)

func Say(names []string) string {
	if len(names) > 0 {
		return fmt.Sprintf("Hello, %s!", strings.Join(names, ", "))
	}

	return "Hello, World!"
}
