package api

import (
	"encoding/json"
	"net/http"

	"github.com/fchazal/noteworthy/server/storage"
	"github.com/fchazal/noteworthy/server/types"
	uuid "github.com/satori/go.uuid"
)

type Note struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

func toNote(n *types.Note) *Note {
	if n == nil {
		return nil
	}
	result := Note{n.Id, n.Name, n.ParentId}

	return &result
}

func noteHandler(w http.ResponseWriter, req *http.Request) {
	var result *Note
	data := []byte("null")

	if id := parseRequest(req, Paths["note"]); id != "" {
		s := storage.Store

		switch req.Method {
		case "GET":
			if s.Notes[id] != nil {
				n := s.Notes[id]
				result = toNote(n)
			}
		case "POST":
			if id == "new" {
				n := &types.Note{}
				json.NewDecoder(req.Body).Decode(&n)

				if s.Books[n.ParentId] == nil && s.Chapters[n.ParentId] == nil {
					n = nil
				} else {
					n.Id = uuid.NewV4().String()
					s.Notes[n.Id] = n

					if s.Books[n.ParentId] != nil {
						s.Books[n.ParentId].AddNote(n.Id)
					}

					if s.Chapters[n.ParentId] != nil {
						s.Chapters[n.ParentId].AddNote(n.Id)
					}

					s.Save()
				}

				result = toNote(n)
			}
		case "PATCH":
			if s.Notes[id] != nil {
				n := s.Notes[id]
				oldParent := n.ParentId
				json.NewDecoder(req.Body).Decode(&n)

				if s.Books[n.ParentId] == nil && s.Chapters[n.ParentId] == nil {
					n.ParentId = oldParent
				}

				if n.ParentId != oldParent {
					if s.Books[n.ParentId] != nil {
						s.Books[n.ParentId].AddNote(n.Id)
					}

					if s.Chapters[n.ParentId] != nil {
						s.Chapters[n.ParentId].AddNote(n.Id)
					}

					if s.Books[oldParent] != nil {
						s.Books[oldParent].RemoveNote(n.Id)
					}

					if s.Chapters[oldParent] != nil {
						s.Chapters[oldParent].RemoveNote(n.Id)
					}
				}
				s.Save()

				result = toNote(n)
			}
		case "DELETE":
			// TODO: what about contained notes and resources
			if s.Notes[id] != nil {
				n := s.Notes[id]
				result = toNote(n)

				delete(s.Notes, n.Id)
				if s.Books[n.ParentId] != nil {
					s.Books[n.ParentId].RemoveNote(n.Id)
				}

				if s.Chapters[n.ParentId] != nil {
					s.Chapters[n.ParentId].RemoveNote(n.Id)
				}

				s.Save()
			}
		}

		data, _ = json.Marshal(result)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
