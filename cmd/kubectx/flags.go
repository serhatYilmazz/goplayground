package main

import "strings"

type Op interface{}

// HelpOp describes printing helo.
type HelpOp struct{}

// ListOp describes listing contexts.
type ListOp struct{}

// SwitchOp indicates intention to switch contexts.
type SwitchOp struct {
	Target string
}

// UnknownOp indicates an unsupported flag.
type UnknownOp struct {
	Args []string
}

// parseArgs looks at flags
// and decides which operation should be taken
func parseArgs(args []string) Op {
	if len(args) == 0 {
		return ListOp{}
	}

	if len(args) == 1 {
		v := args[0]
		if args[0] == "--help" || args[0] == "-h" {
			return HelpOp{}
		}
		if strings.HasPrefix(v, "-") && v != "-" {
			return UnknownOp{Args: args}
		}

		return SwitchOp{Target: args[0]}
	}

	return UnknownOp{}
}
