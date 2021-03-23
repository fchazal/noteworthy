package types

import "time"

// STORE ///////////////////////////////////////////////////////////////////////

type Store struct {
	Path string `json:"path"`
	LibraryContainer
	BookContainer
}

// BOOKSHELF ///////////////////////////////////////////////////////////////////

type Library struct {
	Item
	Path string `json:"path"`
	BookContainer
}

// BOOK ////////////////////////////////////////////////////////////////////////

type Book struct {
	Item
	Path     string `json:"path"`
	ParentId string `json:"parent_id,omitempty"`
	ChapterContainer
	NoteContainer
	ResourceContainer
}

// CHAPTER /////////////////////////////////////////////////////////////////////

type Chapter struct {
	Item
	Path     string `json:"path"`
	ParentId string `json:"parent_id"`
	NoteContainer
	ResourceContainer
}

// NOTE ////////////////////////////////////////////////////////////////////////

type Note struct {
	Item
	Path     string    `json:"path"`
	ParentId string    `json:"parent_id"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// RESOURCE ////////////////////////////////////////////////////////////////////

type Resource struct {
	Item
	ParentId string `json:"parent_id`
	Path     string `json:"path"`
}

// ITEM ////////////////////////////////////////////////////////////////////////

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
