package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fchazal/noteworthy/server/types"
)

type Storage struct {
	Path      string                    `json:"-"`
	Libraries map[string]*types.Library `json:"libraries"`
	Books     map[string]*types.Book    `json:"books"`
	BookStore types.Store               `json:"store"`
}

var Store *Storage

func Open(path string) *Storage {
	if _, err := os.Stat(path); err != nil {
		Store = &Storage{
			path,
			map[string]*types.Library{},
			map[string]*types.Book{},
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
