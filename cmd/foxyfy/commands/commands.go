package commands

import (
	"fmt"
	"io"
)

type Cmd interface {
	Execute(args []string)
}

type RootCmd struct {
	out       io.Writer
	helpCmd   Cmd
	lsCmd     Cmd
	updateCmd Cmd
}

func NewRootCmd(out io.Writer, helpCmd Cmd, lsCmd Cmd, updateCmd Cmd) RootCmd {
	return RootCmd{
		out:       out,
		helpCmd:   helpCmd,
		lsCmd:     lsCmd,
		updateCmd: updateCmd,
	}
}

func (c RootCmd) Execute(args []string) {
	if len(args) < 1 {
		c.helpCmd.Execute([]string{})
		return
	}
	cmd := args[0]
	cmdArgs := args[1:]

	if cmd == "ls" {
		c.lsCmd.Execute(cmdArgs)
	} else if cmd == "help" {
		c.helpCmd.Execute(cmdArgs)
	} else if cmd == "update" {
		c.updateCmd.Execute(cmdArgs)
	} else {
		fmt.Fprintf(c.out, "Unknown command '%s'\n", cmd)
	}
}
