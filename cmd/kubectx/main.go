package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var op Op
	op = parseArgs(os.Args[1:])

	switch v:= op.(type) {
	case UnknownOp:
		fmt.Printf("error: Unsupported Operation: %s\n", strings.Join(v.Args, " "))
		//TODO print --help string
		os.Exit(1)
	case ListOp:
		panic("Not implemented")
	case SwitchOp:
		panic("Not implemented")
	default:
		fmt.Printf("Internal error: Operation type %T not handled\n", op)
	}
}

