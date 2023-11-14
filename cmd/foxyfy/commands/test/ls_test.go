package commands_test

import (
	"bytes"
	"flag"
	"foxyfy/cmd/foxyfy/commands"
	"os"
	"testing"

	"database/sql"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestLs(t *testing.T) {
	file := "./new.db"
	db := createTestDb(file, t)
	defer db.Close()
	defer removeTestDb(file, t)

	var out bytes.Buffer
	rootCmd := commands.NewRootCmd(&out)
	os.Args = []string{"bin", "ls", "-b", "music", "-db", file}
	rootCmd.Execute()

	expected := "song1 - music/song1\n" +
		"---- subdir ----\n" +
		"song2 - music/song2\n" +
		"song3 - music/song3\n"

	assert.Equal(t, expected, out.String())
}

func TestLsWithMissingParams(t *testing.T) {
	args := []string{"bin", "ls", "-db", "./file.db"}
	expected := "Missing -b argument\n"
	testCmd(args, expected, t)

	args = []string{"bin", "ls", "-b", "", "-db", "./file.db"}
	expected = "Missing -b argument\n"
	testCmd(args, expected, t)

	args = []string{"bin", "ls", "-b", "music"}
	expected = "Missing -db argument\n"
	testCmd(args, expected, t)

	args = []string{"bin", "ls", "-b", "music", "-db", ""}
	expected = "Missing -db argument\n"
	testCmd(args, expected, t)
}

func TestLsWithMissingDbFile(t *testing.T) {
	args := []string{"bin", "ls", "-b", "music", "-db", "./file.db"}
	expected := "Database './file.db' does not exist\n"
	testCmd(args, expected, t)
}

func testCmd(args []string, expectedOut string, t *testing.T) {
	var out bytes.Buffer
	rootCmd := commands.NewRootCmd(&out)
	os.Args = args

	resetFlags()
	rootCmd.Execute()

	assert.Equal(t, expectedOut, out.String())
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func createTestDb(file string, t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", file)
	assert.Nil(t, err)

	_, err = db.Exec(
		"drop table if exists moz_bookmarks;" +

			"create table moz_bookmarks (" +
			"id INTEGER PRIMARY KEY," +
			"type INTEGER," +
			"title TEXT," +
			"parent INTEGER," +
			"fk INTEGER);")
	assert.Nil(t, err)

	_, err = db.Exec(
		"drop table if exists moz_places;" +

			"create table moz_places (" +
			"id INTEGER PRIMARY KEY," +
			"url TEXT);")
	assert.Nil(t, err)

	_, err = db.Exec(
		"insert into moz_bookmarks values(1, 2, 'music', null, null);" +
			"insert into moz_bookmarks values(2, 1, 'song1', 1, 1);" +
			"insert into moz_bookmarks values(3, 2, 'subdir', 1, null);" +
			"insert into moz_bookmarks values(4, 1, 'song2', 3, 2);" +
			"insert into moz_bookmarks values(5, 1, 'song3', 3, 3);")
	assert.Nil(t, err)

	_, err = db.Exec(
		"insert into moz_places values(1, 'music/song1');" +
			"insert into moz_places values(2, 'music/song2');" +
			"insert into moz_places values(3, 'music/song3');")
	assert.Nil(t, err)

	return db
}

func removeTestDb(file string, t *testing.T) {
	err := os.Remove(file)
	assert.Nil(t, err)
}
