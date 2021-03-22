package api

import (
	"encoding/json"
	"net/http"

	"github.com/fchazal/noteworthy/server/storage"
	"github.com/fchazal/noteworthy/server/types"
	uuid "github.com/satori/go.uuid"
)

type Book struct {
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	ParentId string       `json:"parent_id"`
	Chapters []types.Item `json:"chapters"`
	Notes    []types.Item `json:"notes"`
}

func toBook(b *types.Book) *Book {
	result := Book{b.Id, b.Name, b.ParentId, []types.Item{}, []types.Item{}}
	/*
		for _, id := range b.Chapters {
			result.Chapters = append(result.Chapters, storage.Store.Chapters[id].Item)
		}

		for _, id := range b.Notes {
			result.Notes = append(result.Notes, storage.Store.Notes[id].Item)
		}
	*/
	return &result
}

func bookHandler(w http.ResponseWriter, req *http.Request) {
	var result *Book
	data := []byte("null")

	if id := parseRequest(req, Paths["book"]); id != "" {
		s := storage.Store

		switch req.Method {
		case "GET":
			if s.Books[id] != nil {
				b := s.Books[id]
				result = toBook(b)
			}
		case "POST":
			if id == "new" {
				b := &types.Book{}
				json.NewDecoder(req.Body).Decode(&b)
				b.Id = uuid.NewV4().String()

				s.Books[b.Id] = b
				if b.ParentId == "" || s.Libraries[b.ParentId] == nil {
					b.ParentId = ""
					s.BookStore.AddBook(b.Id)
				} else {
					s.Libraries[b.ParentId].AddBook(b.Id)
				}
				s.Save()

				result = toBook(b)
			}
		case "PATCH":
			if s.Books[id] != nil {
				b := s.Books[id]
				oldParent := b.ParentId
				json.NewDecoder(req.Body).Decode(&b)

				if s.Libraries[b.ParentId] == nil {
					b.ParentId = ""
				}

				if b.ParentId != oldParent {
					if b.ParentId != "" {
						s.Libraries[b.ParentId].AddBook(b.Id)
					} else {
						s.BookStore.AddBook(b.Id)
					}

					if oldParent != "" {
						s.Libraries[oldParent].RemoveBook(b.Id)
					} else {
						s.BookStore.RemoveBook(b.Id)
					}
				}
				s.Save()

				result = toBook(b)
			}
		case "DELETE":
			if s.Books[id] != nil {
				b := s.Books[id]
				result = toBook(b)

				delete(s.Books, b.Id)
				if b.ParentId == "" {
					s.BookStore.RemoveBook(b.Id)
				} else {
					s.Libraries[b.ParentId].RemoveBook(b.Id)
				}
				s.Save()
			}
		}

		data, _ = json.Marshal(result)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
