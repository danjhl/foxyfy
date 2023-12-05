package commands

import (
	"fmt"
	"io"
)

type HelpCmd struct {
	out io.Writer
}

func NewHelpCmd(out io.Writer) HelpCmd {
	return HelpCmd{out: out}
}

func (c HelpCmd) Execute(args []string) {
	help := "\n" +
		"Usage: [CMD] [FLAGS]\n\n" +
		"ls -b [bookmark]                     List youtube bookmarks in parent [bookmark]\n\n"

	fmt.Fprint(c.out, help)
}
