package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fchazal/noteworthy/server/types"
)

type Storage struct {
	Path      string                     `json:"-"`
	Libraries map[string]*types.Library  `json:"libraries"`
	Books     map[string]*types.Book     `json:"books"`
	Chapters  map[string]*types.Chapter  `json:"chapters"`
	Notes     map[string]*types.Note     `json:"notes"`
	Resources map[string]*types.Resource `json:"resources"`
	BookStore types.Store                `json:"store"`
}

var Store *Storage

func Open(path string) *Storage {
	if _, err := os.Stat(path); err != nil {
		Store = &Storage{
			path,
			map[string]*types.Library{},
			map[string]*types.Book{},
			map[string]*types.Chapter{},
			map[string]*types.Note{},
			map[string]*types.Resource{},
			types.Store{},
		}
	} else {
		if data, err := ioutil.ReadFile(path); err != nil {
			log.Fatal("can't read database")
		} else {
			Store = &Storage{Path: path}
			json.Unmarshal(data, Store)
		}
	}

	return Store
}

func (s *Storage) Save() {
	data, _ := json.MarshalIndent(s, "", "	")

	if err := ioutil.WriteFile(s.Path, data, 0644); err != nil {
		log.Fatal("can't write database")
	}
}

func Save() {
	Store.Save()
}
