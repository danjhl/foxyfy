package commands

import (
	"errors"
	"flag"
	"fmt"
	"foxyfy/internal/bookmarks"
	"io"
	"os"
	"strings"
)

type LsCmd struct {
	out io.Writer
}

func NewLsCmd(out io.Writer) LsCmd {
	return LsCmd{out: out}
}

func (c LsCmd) Execute(args []string) {
	bFlag := flag.String("b", "", "Parent bookmark")
	dbFlag := flag.String("db", "", "Absolute file path to firefox places db")

	flag.CommandLine.Parse(args)

	if *bFlag == "" {
		c.print("Missing -b argument\n")
		return
	}

	if *dbFlag == "" {
		c.print("Missing -db argument\n")
		return
	}

	_, err := os.Stat(*dbFlag)

	if errors.Is(err, os.ErrNotExist) {
		c.print(fmt.Sprintf("Database '%s' does not exist\n", *dbFlag))
		return
	}

	ffdb := bookmarks.FirefoxDb{Source: *dbFlag}

	id, err := ffdb.QueryIdFor(*bFlag)
	if err != nil {
		c.printErr(err)
		return
	}

	bms, err := bookmarks.GetChildBookmarksFor(id[0], "", ffdb)
	if err != nil {
		c.printErr(err)
		return
	}

	var sb strings.Builder

	for _, bm := range bms {
		switch bm.Bm.Bm_type {
		case bookmarks.DirectoryBookmark:

		case bookmarks.UrlBookmark:
			sb.WriteString(fmt.Sprintf("%s/%s (%s)\n", bm.Dir, bm.Bm.Title, bm.Bm.Url))
		}
	}

	c.print(sb.String())
}

func (c LsCmd) print(a ...any) {
	fmt.Fprint(c.out, a...)
}

func (c LsCmd) printErr(err error) {
	fmt.Fprintf(c.out, "Error: %s", err)
}
