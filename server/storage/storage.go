package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fchazal/noteworthy/server/books"
	"github.com/fchazal/noteworthy/server/bookshelves"
	"github.com/fchazal/noteworthy/server/chapters"
	"github.com/fchazal/noteworthy/server/notes"
	"github.com/fchazal/noteworthy/server/resources"
)

type Storage struct {
	Path        string                            `json:"-"`
	Bookshelves map[string]*bookshelves.Bookshelf `json:"bookshelves"`
	Books       map[string]*books.Book            `json:"books"`
	Chapters    map[string]*chapters.Chapter      `json:"chapters"`
	Notes       map[string]*notes.Note            `json:"notes"`
	Resources   map[string]*resources.Resource    `json:"resources"`
	bookshelves.BookshelfContainer
}

var Library *Storage

func Open(path string) {
	if _, err := os.Stat(path); err != nil {
		Library = &Storage{
			path,
			map[string]*bookshelves.Bookshelf{},
			map[string]*books.Book{},
			map[string]*chapters.Chapter{},
			map[string]*notes.Note{},
			map[string]*resources.Resource{},
			bookshelves.NewContainer(),
		}
	} else {
		if data, err := ioutil.ReadFile(path); err != nil {
			log.Fatal("can't read database")
		} else {
			Library = &Storage{Path: path}
			json.Unmarshal(data, Library)
		}
	}

	bookshelves.Bookshelves = &Library.Bookshelves
	books.Books = &Library.Books
	chapters.Chapters = &Library.Chapters
	notes.Notes = &Library.Notes
	resources.Resources = &Library.Resources
}

func Save() {
	data, _ := json.MarshalIndent(Library, "", "	")

	if err := ioutil.WriteFile(Library.Path, data, 0644); err != nil {
		log.Fatal("can't write database")
	}
}
