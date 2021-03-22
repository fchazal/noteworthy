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

/*
// CHAPTER CONTAINER ///////////////////////////////////////////////////////////

type ChapterContainer struct {
	Chapters []string `json:"chapters"`
}

func (c *ChapterContainer) AddChapter(item *Chapter) {
	c.Chapters = append(c.Chapters, item.Id)
}

func (c *ChapterContainer) RemoveChapter(item *Chapter) {
	c.Chapters = utils.RemoveItem(c.Chapters, item.Id)
}

// NOTE CONTAINER //////////////////////////////////////////////////////////////

type NoteContainer struct {
	Notes []string `json:"notes"`
}

func (c *NoteContainer) AddChapter(item *Note) {
	c.Notes = append(c.Notes, item.Id)
}

func (c *NoteContainer) RemoveChapter(item *Note) {
	c.Notes = utils.RemoveItem(c.Notes, item.Id)
}

// RESOURCE CONTAINER //////////////////////////////////////////////////////////

type ResourceContainer struct {
	Resources []string `json:"resource"`
}

func (c *ResourceContainer) AddResourc(item *Resource) {
	c.Resources = append(c.Resources, item.Id)
}

func (c *ResourceContainer) RemoveResource(item *Resource) {
	c.Resources = utils.RemoveItem(c.Resources, item.Id)
}
*/
