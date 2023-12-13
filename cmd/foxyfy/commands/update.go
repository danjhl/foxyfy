package commands

import (
	"flag"
	"fmt"
	"foxyfy/internal/bookmarks"
	"foxyfy/internal/executor"
	"foxyfy/internal/input"
	"io"
	"os"
)

type UpdateCmd struct {
	out      io.Writer
	executor executor.Executor
}

func NewUpdateCmd(out io.Writer, executor executor.Executor) UpdateCmd {
	return UpdateCmd{out: out, executor: executor}
}

func (c UpdateCmd) Name() string {
	return "update"
}

func (c UpdateCmd) Help() string {
	return "update\n\n" +

		"    usage: update -b <bookmark> -db <database> [OPTS]\n\n" +

		"    updates music from youtube bookmarks.\n\n" +

		"    bookmark:    name of directory bookmark to search for youtube bookmarks\n" +
		"    database:    absolute file path to firefox places.sqlite file database\n\n" +

		"    OPTS:\n" +
		"    -o <dir>      output directory for the music updates, default is current directory\n" +
		"    -n            dry-run option which does not download bu displays the updates"
}

func (c UpdateCmd) Execute(args []string) {
	bFlag := flag.String("b", "", "Parent bookmark")
	dbFlag := flag.String("db", "", "Absolute file path to firefox places db")
	oFlag := flag.String("o", ".", "Output directory")
	nFlag := flag.Bool("n", false, "Dry run")

	flag.CommandLine.Parse(args)
	ffdb := bookmarks.FirefoxDb{Source: *dbFlag}

	id, err := ffdb.QueryIdFor(*bFlag)
	if err != nil {
		fmt.Fprintf(c.out, "Error: %s\n", err)
		return
	}

	if len(id) < 1 {
		fmt.Fprintf(c.out, "No bookmark '%s' found\n", *bFlag)
		return
	}

	cbms, err := bookmarks.GetChildBookmarksFor(id[0], "", ffdb)
	if err != nil {
		fmt.Fprintf(c.out, "Error: %s\n", err)
		return
	}

	if *nFlag {
		fmt.Fprint(c.out, "Updates:\n\n")
	}

	for _, cbm := range cbms {
		ok := input.IsValidYoutubeUrl(cbm.Bm.Url)
		if !ok {
			fmt.Fprintf(c.out, "Error: '%s' is not a valid url\n", cbm.Bm.Url)
			continue
		}

		var bmDir string
		if cbm.Dir != "" {
			bmDir = cbm.Dir + "/"
		}

		bmDir = input.SanitizeDirectory(bmDir)
		sanitizedFileName := input.SanitizeFileName(cbm.Bm.Title)
		mp3 := *oFlag + "/" + bmDir + sanitizedFileName + ".mp3"

		_, err := os.Stat(mp3)
		if err == nil {
			fmt.Fprintln(c.out, mp3, "already downloaded")
			continue
		}

		if *nFlag {
			fmt.Fprintln(c.out, mp3)
			continue
		}

		fmt.Fprintln(c.out, "downloading", mp3, "...")
		out := *oFlag + "/" + bmDir + sanitizedFileName + ".%(ext)s"

		err = c.executor.RunFromPath("yt-dlp", "-o", out, "--extract-audio", "--audio-format", "mp3", cbm.Bm.Url)

		if err != nil {
			fmt.Fprintf(c.out, "Error: %s\n", err)
		}

	}
}
