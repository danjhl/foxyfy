package commands_test

import (
	"bytes"
	"flag"
	"fmt"
	"foxyfy/cmd/foxyfy/commands"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpDir string

func TestMain(m *testing.M) {
	tmp, err := os.MkdirTemp("", "tmp")
	if err != nil {
		log.Fatal(err)
	}
	tmpDir = tmp

	file := tmpDir + "/new.db"
	createTestDb(file)
	defer os.RemoveAll(tmpDir)

	code := m.Run()
	os.Exit(code)
}

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
	args := []string{"-b", "music", "-o", tmpDir + "/out", "-db", tmpDir + "/new.db"}

	resetFlags()
	updateCmd.Execute(args)

	expected := []string{
		fmt.Sprintf("yt-dlp -o %s/out/song1.%%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song1", tmpDir),
		fmt.Sprintf("yt-dlp -o %s/out/subdir/song2.%%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song2", tmpDir),
		fmt.Sprintf("yt-dlp -o %s/out/subdir/song3.%%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song3", tmpDir),
		fmt.Sprintf("yt-dlp -o %s/out/subdir/subdir/song4.%%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song4", tmpDir),
	}
	assert.Equal(t, expected, executor.executed)
}

func TestUpdateWithMissingPassedBookmark(t *testing.T) {
	var out bytes.Buffer
	executor := mockExecutor{}
	updateCmd := commands.NewUpdateCmd(&out, &executor)
	args := []string{"-b", "does not exist", "-o", tmpDir + "/out", "-db", tmpDir + "/new.db"}

	resetFlags()
	updateCmd.Execute(args)

	assert.Equal(t, "No bookmark 'does not exist' found\n", out.String())
}

func TestUpdateDryRun(t *testing.T) {
	var out bytes.Buffer
	executor := mockExecutor{}
	updateCmd := commands.NewUpdateCmd(&out, &executor)
	args := []string{"-b", "music", "-o", tmpDir + "/out", "-db", tmpDir + "/new.db", "-n"}

	resetFlags()
	updateCmd.Execute(args)

	assert.Nil(t, executor.executed)

	expectedPrint := "Updates:\n\n" +
		fmt.Sprintf("%s/out/song1.mp3\n", tmpDir) +
		fmt.Sprintf("%s/out/subdir/song2.mp3\n", tmpDir) +
		fmt.Sprintf("%s/out/subdir/song3.mp3\n", tmpDir) +
		fmt.Sprintf("%s/out/subdir/subdir/song4.mp3\n", tmpDir)

	assert.Equal(t, expectedPrint, out.String())
}

func TestUpdateWithExistingFiles(t *testing.T) {
	var out bytes.Buffer
	executor := mockExecutor{}
	updateCmd := commands.NewUpdateCmd(&out, &executor)
	args := []string{"-b", "music", "-o", tmpDir + "/out", "-db", tmpDir + "/new.db"}

	createSong1Mp3()
	defer os.RemoveAll(tmpDir + "/out")

	resetFlags()
	updateCmd.Execute(args)

	expected := []string{
		fmt.Sprintf("yt-dlp -o %s/out/subdir/song2.%%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song2", tmpDir),
		fmt.Sprintf("yt-dlp -o %s/out/subdir/song3.%%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song3", tmpDir),
		fmt.Sprintf("yt-dlp -o %s/out/subdir/subdir/song4.%%(ext)s --extract-audio --audio-format mp3 https://www.youtube.com/watch?v=song4", tmpDir),
	}
	assert.Equal(t, expected, executor.executed)
}

func TestUpdateWithExistingFilesDryRun(t *testing.T) {
	var out bytes.Buffer
	executor := mockExecutor{}
	updateCmd := commands.NewUpdateCmd(&out, &executor)
	args := []string{"-b", "music", "-o", tmpDir + "/out", "-db", tmpDir + "/new.db", "-n"}

	createSong1Mp3()
	defer os.RemoveAll(tmpDir + "/out")

	resetFlags()
	updateCmd.Execute(args)

	assert.Nil(t, executor.executed)

	expectedPrint := "Updates:\n\n" +
		fmt.Sprintf("%s/out/song1.mp3 already downloaded\n", tmpDir) +
		fmt.Sprintf("%s/out/subdir/song2.mp3\n", tmpDir) +
		fmt.Sprintf("%s/out/subdir/song3.mp3\n", tmpDir) +
		fmt.Sprintf("%s/out/subdir/subdir/song4.mp3\n", tmpDir)

	assert.Equal(t, expectedPrint, out.String())
}

func createSong1Mp3() {
	err := os.Mkdir(tmpDir+"/out", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	mp3, err := os.Create(tmpDir + "/out/song1.mp3")
	if err != nil {
		log.Fatal(err)
	}
	mp3.Close()
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
