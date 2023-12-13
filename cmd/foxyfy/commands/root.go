package commands

import (
	"fmt"
	"io"
	"strings"
)

type RootCmd struct {
	out  io.Writer
	cmds []Cmd
}

func NewRootCmd(out io.Writer, cmds []Cmd) RootCmd {
	return RootCmd{
		out:  out,
		cmds: cmds,
	}
}

func (c RootCmd) Execute(args []string) {
	if len(args) < 1 {
		c.help()
		return
	}
	command := args[0]
	commandArgs := args[1:]

	if command == "help" {
		c.help()
		return
	}

	for _, cmd := range c.cmds {
		if cmd.Name() == command {
			cmd.Execute(commandArgs)
			return
		}
	}

	fmt.Fprintf(c.out, "Unknown command '%s'\n", command)
}

func (c RootCmd) help() {
	var sb strings.Builder

	for _, cmd := range c.cmds {
		sb.WriteString(cmd.Help())
		sb.WriteString("\n\n")
	}

	fmt.Fprint(c.out, sb.String())
}
