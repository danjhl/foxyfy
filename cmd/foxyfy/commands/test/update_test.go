package commands_test

import (
	"bytes"
	"foxyfy/cmd/foxyfy/commands"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockExecutor struct {
	executed []string
}

func (d *mockExecutor) RunFromPath(cmd string, args ...string) error {
	d.executed = append(d.executed, cmd+" "+strings.Join(args, " "))
	return nil
}

func TestUpdate(t *testing.T) {
	var out bytes.Buffer
	executor := mockExecutor{}
	updateCmd := commands.NewUpdateCmd(&out, &executor)
	args := []string{"-b", "music", "-o", "./out", "-db", "./new.db"}

	file := "./new.db"
	db := createTestDb(file, t)
	defer db.Close()
	defer removeTestDb(file, t)

	resetFlags()
	updateCmd.Execute(args)

	expected := []string{
		"yt-dlp -o ./out/song1.%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song1",
		"yt-dlp -o ./out/subdir/song2.%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song2",
		"yt-dlp -o ./out/subdir/song3.%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song3",
		"yt-dlp -o ./out/subdir/subdir/song4.%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song4",
	}
	assert.Equal(t, expected, executor.executed)
}

func TestUpdateWithMissingPassedBookmark(t *testing.T) {
	var out bytes.Buffer
	executor := mockExecutor{}
	updateCmd := commands.NewUpdateCmd(&out, &executor)
	args := []string{"-b", "does not exist", "-o", "./out", "-db", "./new.db"}

	file := "./new.db"
	db := createTestDb(file, t)
	defer db.Close()
	defer removeTestDb(file, t)

	resetFlags()
	updateCmd.Execute(args)

	assert.Equal(t, "No bookmark 'does not exist' found\n", out.String())
}
