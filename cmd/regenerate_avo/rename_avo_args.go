package main

import (
	"fmt"
	"strings"
)

func renameAvoArgs(args string, sizeClass int) string {
	parts := strings.Split(args, ", ")
	renamedParts := []string{}
	for i, part := range parts {
		if i < len(parts)-1 {
			renamedParts = append(renamedParts, renameArg(part, sizeClass))
		} else {
			// Replace last arg with retX/Y/Z arg
			renamedParts = append(renamedParts, returnArg(sizeClass))
		}
	}
	return strings.Join(renamedParts, ", ")
}

func renameArg(arg string, sizeClass int) string {
	// We _expect_ to need a number of distinct renaming cases
	// Add them there
	switch {
	case strings.HasPrefix(arg, "reg"):
		naked := strings.ToLower(strings.TrimPrefix(arg, "reg"))
		return naked + registerSuffix(sizeClass)
	default:
		return arg
	}
}

func returnArg(sizeClass int) string {
	return "ret" + registerSuffix(sizeClass)
}

// NB: This relies on the register naming scheme used by the code generator not changing.
// There's probably a better way to manage this.
func registerSuffix(sizeClass int) string {
	switch sizeClass {
	case 128:
		return "X"
	case 256:
		return "Y"
	case 512:
		return "Z"
	default:
		panic(fmt.Errorf("unsupported size-class %d", sizeClass))
	}
}
