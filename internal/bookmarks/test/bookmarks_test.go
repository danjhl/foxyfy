package bookmarks

import (
	"foxyfy/internal/bookmarks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Bookmark = bookmarks.Bookmark

type DummyDb struct {
	ids       []int
	bookmarks []Bookmark
}

func (db DummyDb) QueryIdFor(bookmark string) ([]int, error) {
	var result []int
	for _, bm := range db.bookmarks {
		if bm.Title == bookmark {
			result = append(result, bm.Id)
		}
	}
	return result, nil
}

func (db DummyDb) QueryChildrenFor(id int) ([]Bookmark, error) {
	var result []Bookmark
	for _, bm := range db.bookmarks {
		if bm.Parent == id {
			result = append(result, bm)
		}
	}
	return result, nil
}

func TestGetIdFor(t *testing.T) {
	db := DummyDb{
		bookmarks: []Bookmark{
			Bookmark{Id: 1, Parent: -1, Bm_type: bookmarks.DirectoryBookmark, Title: "music"},
		},
	}

	result, err := bookmarks.GetIdFor("music", db)
	assert.Nil(t, err)
	assert.Equal(t, []int{1}, result)
}

func TestQueryChildrenFor(t *testing.T) {
	db := DummyDb{bookmarks: []Bookmark{
		Bookmark{Id: 10, Parent: 1, Bm_type: bookmarks.DirectoryBookmark, Title: "child"},
	}}

	result, err := bookmarks.GetChildrenFor(1, db)
	assert.Nil(t, err)

	expected := []Bookmark{
		Bookmark{Id: 10, Parent: 1, Bm_type: bookmarks.DirectoryBookmark, Title: "child"},
	}
	assert.Equal(t, expected, result)
}

func TestGetFlatBookmarksFor(t *testing.T) {
	db := DummyDb{
		bookmarks: []Bookmark{
			Bookmark{Id: 4, Parent: 3, Bm_type: bookmarks.UrlBookmark, Title: "music.io/song2"},
			Bookmark{Id: 5, Parent: 10, Bm_type: bookmarks.UrlBookmark, Title: "music.io/song3"},
			Bookmark{Id: 2, Parent: 1, Bm_type: bookmarks.UrlBookmark, Title: "music.io/song1"},
			Bookmark{Id: 3, Parent: 1, Bm_type: bookmarks.DirectoryBookmark, Title: "favorites"},
		},
	}

	flat, err := bookmarks.GetFlatBookmarksFor(1, db)
	assert.Nil(t, err)

	expected := []Bookmark{
		Bookmark{Id: 2, Parent: 1, Bm_type: bookmarks.UrlBookmark, Title: "music.io/song1"},
		Bookmark{Id: 3, Parent: 1, Bm_type: bookmarks.DirectoryBookmark, Title: "favorites"},
		Bookmark{Id: 4, Parent: 3, Bm_type: bookmarks.UrlBookmark, Title: "music.io/song2"},
	}
	assert.Equal(t, expected, flat)
}
