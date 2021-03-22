package types

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
	//	NoteContainer
	//	ChapterContainer
	//	ResourceContainer
}

/*
// CHAPTER /////////////////////////////////////////////////////////////////////

type Chapter struct {
	Item
	Path   string `json:"path"`
	Parent string `json:"parent_id"`
	NoteContainer
	ResourceContainer
}

// NOTE ////////////////////////////////////////////////////////////////////////

type Note struct {
	Item
	Path    string `json:"-"`
	Parent  string `json:"parent_id"`
	Created time.Time
	Updated time.Time
	ResourceContainer
}

// RESOURCE ////////////////////////////////////////////////////////////////////

type Resource struct {
	Item
	Path string
}
*/
// ITEM ////////////////////////////////////////////////////////////////////////

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
