package commands

import (
	"flag"
	"fmt"
	"foxyfy/internal/bookmarks"
	"foxyfy/internal/input"
	"io"

	"errors"
)

type Executor interface {
	RunFromPath(cmd string, args ...string) error
}

type UpdateCmd struct {
	out      io.Writer
	executor Executor
}

func NewUpdateCmd(out io.Writer, executor Executor) UpdateCmd {
	return UpdateCmd{out: out, executor: executor}
}

func (c UpdateCmd) Execute(args []string) {
	bFlag := flag.String("b", "", "Parent bookmark")
	dbFlag := flag.String("db", "", "Absolute file path to firefox places db")
	oFlag := flag.String("o", ".", "Output directory")

	flag.CommandLine.Parse(args)
	ffdb := bookmarks.FirefoxDb{Source: *dbFlag}

	id, err := ffdb.QueryIdFor(*bFlag)
	if err != nil {
		c.printErr(err)
		return
	}

	if len(id) < 1 {
		c.print(fmt.Sprintf("No bookmark '%s' found\n", *bFlag))
		return
	}

	cbms, err := bookmarks.GetChildBookmarksFor(id[0], "", ffdb)
	if err != nil {
		c.printErr(err)
		return
	}

	for _, cbm := range cbms {
		ok := input.IsValidYoutubeUrl(cbm.Bm.Url)
		if !ok {
			c.printErr(errors.New(fmt.Sprintf("'%s' is not a valid url", cbm.Bm.Url)))
			continue
		}

		var bmDir string
		if cbm.Dir != "" {
			bmDir = cbm.Dir + "/"
		}

		bmDir = input.SanitizeDirectory(bmDir)
		sanitizedFileName := input.SanitizeFileName(cbm.Bm.Title)

		err := c.executor.RunFromPath("yt-dlp",
			"-o",
			*oFlag+"/"+bmDir+sanitizedFileName+".%(ext)s",
			"--extract-audio",
			"--audio-format",
			"mp3",
			cbm.Bm.Url)

		if err != nil {
			c.printErr(err)
		}
	}
}

func (c UpdateCmd) print(a ...any) {
	fmt.Fprint(c.out, a...)
}

func (c UpdateCmd) printErr(err error) {
	fmt.Fprintf(c.out, "Error: %s", err)
}
