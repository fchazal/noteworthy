package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fchazal/noteworthy/server/api/books"
	"github.com/fchazal/noteworthy/server/api/chapters"
	"github.com/fchazal/noteworthy/server/api/libraries"
	"github.com/fchazal/noteworthy/server/api/notes"
	"github.com/fchazal/noteworthy/server/api/resources"
)

type Storage struct {
	Path      string                         `json:"-"`
	Libraries map[string]*libraries.Library  `json:"libraries"`
	Books     map[string]*books.Book         `json:"books"`
	Chapters  map[string]*chapters.Chapter   `json:"chapters"`
	Notes     map[string]*notes.Note         `json:"notes"`
	Resources map[string]*resources.Resource `json:"resources"`
	libraries.LibraryContainer
}

var Library *Storage

func Open(path string) *Storage {
	if _, err := os.Stat(path); err != nil {
		Library = &Storage{
			path,
			map[string]*libraries.Library{},
			map[string]*books.Book{},
			map[string]*chapters.Chapter{},
			map[string]*notes.Note{},
			map[string]*resources.Resource{},
			libraries.NewContainer(),
		}
	} else {
		if data, err := ioutil.ReadFile(path); err != nil {
			log.Fatal("can't read database")
		} else {
			Library = &Storage{Path: path}
			json.Unmarshal(data, Library)
		}
	}

	return Library
}

func (s *Storage) Save() {
	data, _ := json.MarshalIndent(s, "", "	")

	if err := ioutil.WriteFile(s.Path, data, 0644); err != nil {
		log.Fatal("can't write database")
	}
}

func Save() {
	Library.Save()
}
