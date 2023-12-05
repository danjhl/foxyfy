package bookmarks

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Type int

const (
	DirectoryBookmark Type = 2
	UrlBookmark            = 1
)

type Bookmark struct {
	Id      int
	Parent  int
	Bm_type Type
	Title   string
	Url     string
}

type Db interface {
	QueryIdFor(bookmark string) ([]int, error)
	QueryChildrenFor(id int) ([]Bookmark, error)
}

type FirefoxDb struct {
	Source string
}

func (s FirefoxDb) QueryIdFor(bookmark string) ([]int, error) {
	db, err := sql.Open("sqlite", s.Source)
	defer db.Close()

	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select id from moz_bookmarks where title = ?", bookmark)
	if err != nil {
		return nil, err
	} else {
		var id int
		var ids []int
		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				return nil, err
			}
			ids = append(ids, id)
		}
		return ids, nil
	}
}

func (s FirefoxDb) QueryChildrenFor(id int) ([]Bookmark, error) {
	db, err := sql.Open("sqlite", s.Source)
	defer db.Close()

	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select id, parent, type, title, fk from moz_bookmarks where parent = ?", id)
	if err != nil {
		return nil, err
	}

	var bookmarks []Bookmark

	for rows.Next() {
		var bookmark Bookmark
		var fk *int
		err := rows.Scan(&bookmark.Id, &bookmark.Parent, &bookmark.Bm_type, &bookmark.Title, &fk)
		if err != nil {
			return nil, err
		}
		if bookmark.Bm_type == UrlBookmark {
			rows2, err := db.Query("select url from moz_places where id = ?", fk)
			if err != nil {
				return nil, err
			}
			if rows2.Next() {
				err := rows2.Scan(&bookmark.Url)
				if err != nil {
					return nil, err
				}
			}
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

func GetIdFor(bookmark string, db Db) ([]int, error) {
	return db.QueryIdFor(bookmark)
}

func GetChildrenFor(id int, db Db) ([]Bookmark, error) {
	return db.QueryChildrenFor(id)
}

type ChildBookmark struct {
	Dir string
	Bm  Bookmark
}

func GetChildBookmarksFor(id int, parentDir string, db Db) ([]ChildBookmark, error) {
	children, err := GetChildrenFor(id, db)

	if err != nil {
		return nil, err
	}
	var flat []ChildBookmark
	for _, child := range children {
		withDir := ChildBookmark{Dir: parentDir, Bm: child}
		if child.Bm_type == DirectoryBookmark {
			var dir string
			if parentDir != "" {
				dir = parentDir + "/" + child.Title
			} else {
				dir = child.Title
			}

			nested, err := GetChildBookmarksFor(child.Id, dir, db)
			if err != nil {
				return nil, err
			}
			flat = append(flat, nested...)
		} else {
			flat = append(flat, withDir)
		}
	}
	return flat, nil
}
