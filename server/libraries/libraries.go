package libraries

import (
	"github.com/fchazal/noteworthy/server/books"
	"github.com/fchazal/noteworthy/server/utils"
	uuid "github.com/satori/go.uuid"
)

// BOOKSHELVES /////////////////////////////////////////////////////////////////

type Library struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"-"`
	books.BookContainer
}

func New(name string) *Library {
	return &Library{
		uuid.NewV4().String(),
		name,
		utils.ToPath(name),
		books.NewContainer(),
	}
}

// ALL BOOKSHELVES /////////////////////////////////////////////////////////////

var Libraries *map[string]*Library

func addLibrary(b *Library) {
	(*Libraries)[b.Id] = b
}

func removeLibrary(b *Library) {
	delete(*Libraries, b.Id)
}

// BOOKSHELF CONTAINERS ////////////////////////////////////////////////////////

type LibraryContainer struct {
	Libraries []string `json:"libraries"`
}

func NewContainer() LibraryContainer {
	return LibraryContainer{
		[]string{},
	}
}

func (c *LibraryContainer) AddLibrary(s *Library) {
	addLibrary(s)
	c.Libraries = append(c.Libraries, s.Id)
}

func (c *LibraryContainer) RemoveLibrary(s *Library) {
	removeLibrary(s)
	c.Libraries = utils.RemoveItem(c.Libraries, s.Id)
}
