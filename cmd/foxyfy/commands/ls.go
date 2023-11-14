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

func (c LsCmd) Execute() {
	bFlag := flag.String("b", "", "Parent bookmark")
	dbFlag := flag.String("db", "", "Absolute file path to firefox places db")

	flag.CommandLine.Parse(os.Args[2:])

	if *bFlag == "" {
		c.fprint("Missing -b argument\n")
		return
	}

	if *dbFlag == "" {
		c.fprint("Missing -db argument\n")
		return
	}

	_, err := os.Stat(*dbFlag)

	if errors.Is(err, os.ErrNotExist) {
		c.fprint("Database '" + *dbFlag + "' does not exist\n")
		return
	}

	ffdb := bookmarks.FirefoxDb{Source: *dbFlag}

	id, err := ffdb.QueryIdFor(*bFlag)
	if err != nil {
		c.fprintErr(err)
		return
	}

	bms, err := bookmarks.GetFlatBookmarksFor(id[0], ffdb)
	if err != nil {
		c.fprintErr(err)
		return
	}

	var sb strings.Builder

	for _, bm := range bms {
		switch bm.Bm_type {
		case bookmarks.DirectoryBookmark:
			sb.WriteString(fmt.Sprintf("---- %s ----\n", bm.Title))
		case bookmarks.UrlBookmark:
			sb.WriteString(fmt.Sprintf("%s - %s\n", bm.Title, bm.Url))
		}
	}

	c.fprint(sb.String())
}

func (c LsCmd) fprint(a ...any) {
	fmt.Fprint(c.out, a...)
}

func (c LsCmd) fprintErr(err error) {
	fmt.Fprintf(c.out, "Error: %s", err)
}
