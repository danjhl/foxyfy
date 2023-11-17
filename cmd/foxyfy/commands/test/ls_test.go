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
	lsCmd := commands.NewLsCmd(&out)
	args := []string{"-b", "music", "-db", file}
	lsCmd.Execute(args)

	expected := "/song1 (https://www.youtube.com/watch?v=song1)\n" +
		"subdir/song2 (https://www.youtube.com/watch?v=song2)\n" +
		"subdir/song3 (https://www.youtube.com/watch?v=song3)\n" +
		"subdir/subdir/song4 (https://www.youtube.com/watch?v=song4)\n"

	assert.Equal(t, expected, out.String())
}

func TestLsWithMissingParams(t *testing.T) {
	args := []string{"-db", "./file.db"}
	expected := "Missing -b argument\n"
	testCmd(args, expected, t)

	args = []string{"-b", "", "-db", "./file.db"}
	expected = "Missing -b argument\n"
	testCmd(args, expected, t)

	args = []string{"-b", "music"}
	expected = "Missing -db argument\n"
	testCmd(args, expected, t)

	args = []string{"-b", "music", "-db", ""}
	expected = "Missing -db argument\n"
	testCmd(args, expected, t)
}

func TestLsWithMissingDbFile(t *testing.T) {
	args := []string{"-b", "music", "-db", "./file.db"}
	expected := "Database './file.db' does not exist\n"
	testCmd(args, expected, t)
}

func testCmd(args []string, expectedOut string, t *testing.T) {
	var out bytes.Buffer
	lsCmd := commands.NewLsCmd(&out)

	resetFlags()
	lsCmd.Execute(args)

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
			"insert into moz_bookmarks values(5, 1, 'song3', 3, 3);" +
			"insert into moz_bookmarks values(6, 2, 'subdir', 3, null);" +
			"insert into moz_bookmarks values(7, 1, 'song4', 6, 4);")
	assert.Nil(t, err)

	_, err = db.Exec(
		"insert into moz_places values(1, 'https://www.youtube.com/watch?v=song1');" +
			"insert into moz_places values(2, 'https://www.youtube.com/watch?v=song2');" +
			"insert into moz_places values(3, 'https://www.youtube.com/watch?v=song3');" +
			"insert into moz_places values(4, 'https://www.youtube.com/watch?v=song4');")
	assert.Nil(t, err)

	return db
}

func removeTestDb(file string, t *testing.T) {
	err := os.Remove(file)
	assert.Nil(t, err)
}
