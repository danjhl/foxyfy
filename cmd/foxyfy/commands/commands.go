package commands

import (
	"fmt"
	"io"
	"os"
)

type RootCmd struct {
	out io.Writer
}

func NewRootCmd(out io.Writer) RootCmd {
	return RootCmd{out: out}
}

func (c RootCmd) Execute() {
	if len(os.Args) < 2 {
		helpCmd := NewHelpCmd(c.out)
		helpCmd.Execute()
		return
	}
	cmd := os.Args[1]

	if cmd == "ls" {
		lsCmd := NewLsCmd(c.out)
		lsCmd.Execute()
	} else if cmd == "help" {
		helpCmd := NewHelpCmd(c.out)
		helpCmd.Execute()
	} else {
		fmt.Fprintf(c.out, "Unknown command '%s'\n", cmd)
	}
}
