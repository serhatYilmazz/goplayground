package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	var op Op
	op = parseArgs(os.Args[1:])

	switch v:= op.(type) {
	case UnknownOp:
		fmt.Printf("%s: Unsupported Operation: %s\n", color.RedString("error"), strings.Join(v.Args, " "))
		//TODO print --help string
		printHelp(os.Stdout)
		os.Exit(1)
	case HelpOp:
		printHelp(os.Stdout)
	case ListOp:
		printListContexts(os.Stdout)
	case SwitchOp:
		panic("Not implemented")
	default:
		fmt.Printf("Internal error: Operation type %T not handled\n", op)
	}
}

