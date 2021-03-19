package books

import (
	"github.com/fchazal/noteworthy/server/chapters"
	"github.com/fchazal/noteworthy/server/notes"
	"github.com/fchazal/noteworthy/server/resources"
	"github.com/fchazal/noteworthy/server/utils"
	uuid "github.com/satori/go.uuid"
)

// BOOKS ///////////////////////////////////////////////////////////////////////

type Book struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent_id"`
	notes.NoteContainer
	chapters.ChapterContainer
	resources.ResourceContainer
}

func New(name string) *Book {
	return &Book{
		uuid.NewV4().String(),
		name,
		"",
		notes.NewContainer(),
		chapters.NewContainer(),
		resources.NewContainer(),
	}
}

// ALL BOOKS ///////////////////////////////////////////////////////////////////

var Items *map[string]*Book

func addBook(b *Book) {
	(*Items)[b.Id] = b
}

func removeBook(b *Book) {
	delete(*Items, b.Id)
}

// BOOK CONTAINERS /////////////////////////////////////////////////////////////

type BookContainer struct {
	Books []string `json:"books"`
}

func NewContainer() BookContainer {
	return BookContainer{
		[]string{},
	}
}
func (c *BookContainer) AddBook(b *Book) {
	addBook(b)
	c.Books = append(c.Books, b.Id)
}

func (c *BookContainer) RemoveBook(b *Book) {
	removeBook(b)
	c.Books = utils.RemoveItem(c.Books, b.Id)
}
