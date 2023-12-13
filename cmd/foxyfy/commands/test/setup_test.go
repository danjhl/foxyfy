package commands_test

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func createTestDb(file string) {
	db, err := sql.Open("sqlite", file)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"drop table if exists moz_bookmarks;" +

			"create table moz_bookmarks (" +
			"id INTEGER PRIMARY KEY," +
			"type INTEGER," +
			"title TEXT," +
			"parent INTEGER," +
			"fk INTEGER);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"drop table if exists moz_places;" +

			"create table moz_places (" +
			"id INTEGER PRIMARY KEY," +
			"url TEXT);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"insert into moz_bookmarks values(1, 2, 'music', null, null);" +
			"insert into moz_bookmarks values(2, 1, 'song1', 1, 1);" +
			"insert into moz_bookmarks values(3, 2, 'subdir', 1, null);" +
			"insert into moz_bookmarks values(4, 1, 'song2', 3, 2);" +
			"insert into moz_bookmarks values(5, 1, 'song3', 3, 3);" +
			"insert into moz_bookmarks values(6, 2, 'subdir', 3, null);" +
			"insert into moz_bookmarks values(7, 1, 'song4', 6, 4);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"insert into moz_places values(1, 'https://www.youtube.com/watch?v=song1');" +
			"insert into moz_places values(2, 'https://www.youtube.com/watch?v=song2');" +
			"insert into moz_places values(3, 'https://www.youtube.com/watch?v=song3');" +
			"insert into moz_places values(4, 'https://www.youtube.com/watch?v=song4');")
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}
