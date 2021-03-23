package api

import (
	"encoding/json"
	"net/http"

	"github.com/fchazal/noteworthy/server/storage"
	"github.com/fchazal/noteworthy/server/types"
	uuid "github.com/satori/go.uuid"
)

type Library struct {
	Id    string       `json:"id"`
	Name  string       `json:"name"`
	Books []types.Item `json:"books"`
}

func toLibrary(l *types.Library) *Library {
	result := Library{l.Id, l.Name, []types.Item{}}

	for _, id := range l.Books {
		result.Books = append(result.Books, storage.Store.Books[id].Item)
	}

	return &result
}

func libraryHandler(w http.ResponseWriter, req *http.Request) {
	var result *Library
	data := []byte("null")

	if id := parseRequest(req, Paths["library"]); id != "" {
		s := storage.Store

		switch req.Method {
		case "GET":
			if s.Libraries[id] != nil {
				l := s.Libraries[id]
				result = toLibrary(l)
			}
		case "POST":
			if id == "new" {
				l := &types.Library{}
				json.NewDecoder(req.Body).Decode(&l)
				l.Id = uuid.NewV4().String()

				s.Libraries[l.Id] = l
				s.BookStore.AddLibrary(l.Id)
				s.Save()

				result = toLibrary(l)
			}
		case "PATCH":
			if s.Libraries[id] != nil {
				l := s.Libraries[id]
				json.NewDecoder(req.Body).Decode(&l)
				s.Save()

				result = toLibrary(l)
			}
		case "DELETE":
			// TODO: what about contained books

			if s.Libraries[id] != nil {
				l := s.Libraries[id]
				result = toLibrary(l)

				delete(s.Libraries, l.Id)
				s.BookStore.RemoveLibrary(l.Id)
				s.Save()
			}
		}

		data, _ = json.Marshal(result)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
