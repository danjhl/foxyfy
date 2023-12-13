package bookmarks

import (
	"foxyfy/internal/bookmarks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Bookmark = bookmarks.Bookmark
type ChildBookmark = bookmarks.ChildBookmark

type MockDb struct {
	ids            []int
	bookmarks      []Bookmark
	childBookmarks []ChildBookmark
}

func (db MockDb) QueryIdFor(bookmark string) ([]int, error) {
	var result []int
	for _, bm := range db.bookmarks {
		if bm.Title == bookmark {
			result = append(result, bm.Id)
		}
	}
	return result, nil
}

func (db MockDb) QueryChildrenFor(id int) ([]Bookmark, error) {
	var result []Bookmark
	for _, bm := range db.bookmarks {
		if bm.Parent == id {
			result = append(result, bm)
		}
	}
	return result, nil
}

func TestGetIdFor(t *testing.T) {
	db := MockDb{
		bookmarks: []Bookmark{
			{Id: 1, Parent: -1, Type: bookmarks.DirectoryBookmark, Title: "music"},
		},
	}

	result, err := bookmarks.GetIdFor("music", db)
	assert.Nil(t, err)
	assert.Equal(t, []int{1}, result)
}

func TestQueryChildrenFor(t *testing.T) {
	db := MockDb{bookmarks: []Bookmark{
		{Id: 10, Parent: 1, Type: bookmarks.DirectoryBookmark, Title: "child"},
	}}

	result, err := bookmarks.GetChildrenFor(1, db)
	assert.Nil(t, err)

	expected := []Bookmark{
		{Id: 10, Parent: 1, Type: bookmarks.DirectoryBookmark, Title: "child"},
	}
	assert.Equal(t, expected, result)
}

func TestGetChildBookmarksFor(t *testing.T) {
	db := MockDb{
		bookmarks: []Bookmark{
			{Id: 4, Parent: 3, Type: bookmarks.UrlBookmark, Title: "music.io/song2"},
			{Id: 5, Parent: 10, Type: bookmarks.UrlBookmark, Title: "music.io/song3"},
			{Id: 2, Parent: 1, Type: bookmarks.UrlBookmark, Title: "music.io/song1"},
			{Id: 3, Parent: 1, Type: bookmarks.DirectoryBookmark, Title: "favorites"},
			{Id: 6, Parent: 3, Type: bookmarks.DirectoryBookmark, Title: "other"},
			{Id: 7, Parent: 6, Type: bookmarks.UrlBookmark, Title: "music.io/song4"},
		},
	}

	flat, err := bookmarks.GetChildBookmarksFor(1, "", db)
	assert.Nil(t, err)

	expected := []ChildBookmark{
		{Dir: "", Bm: Bookmark{Id: 2, Parent: 1, Type: bookmarks.UrlBookmark, Title: "music.io/song1"}},
		{Dir: "favorites", Bm: Bookmark{Id: 4, Parent: 3, Type: bookmarks.UrlBookmark, Title: "music.io/song2"}},
		{Dir: "favorites/other", Bm: Bookmark{Id: 7, Parent: 6, Type: bookmarks.UrlBookmark, Title: "music.io/song4"}},
	}
	assert.Equal(t, expected, flat)
}
