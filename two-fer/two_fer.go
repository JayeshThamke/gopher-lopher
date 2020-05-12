package twofer

import (
	"fmt"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
