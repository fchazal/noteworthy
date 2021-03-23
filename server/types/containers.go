package types

import "github.com/fchazal/noteworthy/server/utils"

// BOOKSHELF CONTAINERS ////////////////////////////////////////////////////////

type LibraryContainer struct {
	Libraries []string `json:"libraries,omitempty"`
}

func (c *LibraryContainer) AddLibrary(id string) {
	if !utils.InSlice(c.Libraries, id) {
		c.Libraries = append(c.Libraries, id)
	}
}

func (c *LibraryContainer) RemoveLibrary(id string) {
	c.Libraries = utils.RemoveFromSlice(c.Libraries, id)
}

// BOOK CONTAINER //////////////////////////////////////////////////////////////

type BookContainer struct {
	Books []string `json:"books,omitempty"`
}

func (c *BookContainer) AddBook(id string) {
	if !utils.InSlice(c.Books, id) {
		c.Books = append(c.Books, id)
	}
}

func (c *BookContainer) RemoveBook(id string) {
	c.Books = utils.RemoveFromSlice(c.Books, id)
}

// CHAPTER CONTAINER ///////////////////////////////////////////////////////////

type ChapterContainer struct {
	Chapters []string `json:"chapters"`
}

func (c *ChapterContainer) AddChapter(id string) {
	if !utils.InSlice(c.Chapters, id) {
		c.Chapters = append(c.Chapters, id)
	}
}

func (c *ChapterContainer) RemoveChapter(id string) {
	c.Chapters = utils.RemoveFromSlice(c.Chapters, id)
}

// NOTE CONTAINER //////////////////////////////////////////////////////////////

type NoteContainer struct {
	Notes []string `json:"notes"`
}

func (c *NoteContainer) AddNote(id string) {
	if !utils.InSlice(c.Notes, id) {
		c.Notes = append(c.Notes, id)
	}
}

func (c *NoteContainer) RemoveNote(id string) {
	c.Notes = utils.RemoveFromSlice(c.Notes, id)
}

// RESOURCE CONTAINER //////////////////////////////////////////////////////////

type ResourceContainer struct {
	Resources []string `json:"resource"`
}

func (c *ResourceContainer) AddResource(id string) {
	if !utils.InSlice(c.Resources, id) {
		c.Resources = append(c.Resources, id)
	}
}

func (c *ResourceContainer) RemoveResource(id string) {
	c.Resources = utils.RemoveFromSlice(c.Resources, id)
}
