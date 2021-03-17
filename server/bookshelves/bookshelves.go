package bookshelves

import (
	"github.com/fchazal/noteworthy/server/books"
	uuid "github.com/satori/go.uuid"
)

// BOOKSHELVES /////////////////////////////////////////////////////////////////

type Bookshelf struct {
	Id   string  `json:"id"`
	Name *string `json:"name"`
	books.BookContainer
}

func New(name string) *Bookshelf {
	return &Bookshelf{
		uuid.NewV4().String(),
		&name,
		books.NewContainer(),
	}
}

// ALL BOOKSHELVES /////////////////////////////////////////////////////////////

var Bookshelves *map[string]*Bookshelf

func addBookshelf(b *Bookshelf) {
	(*Bookshelves)[b.Id] = b
}

func removeBookshelf(b *Bookshelf) {
	delete(*Bookshelves, b.Id)
}

// BOOKSHELF CONTAINERS ////////////////////////////////////////////////////////

type BookshelfContainer struct {
	Bookshelves []string `json:"bookshelves"`
}

func NewContainer() BookshelfContainer {
	return BookshelfContainer{
		[]string{},
	}
}

func (c *BookshelfContainer) AddBookshelf(s *Bookshelf) {
	addBookshelf(s)
	c.Bookshelves = append(c.Bookshelves, s.Id)
}

func (c *BookshelfContainer) RemoveBookshelf(s *Bookshelf) {
	findAndDelete := func(s []string, e string) []string {
		x := 0
		for _, i := range s {
			if i != e {
				s[x] = i
				x++
			}
		}
		return s[:x]
	}
	removeBookshelf(s)
	c.Bookshelves = findAndDelete(c.Bookshelves, s.Id)
}
