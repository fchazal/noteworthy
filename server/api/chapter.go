package api

import (
	"encoding/json"
	"net/http"

	"github.com/fchazal/noteworthy/server/storage"
	"github.com/fchazal/noteworthy/server/types"
	uuid "github.com/satori/go.uuid"
)

type Chapter struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	ParentId  string       `json:"parent_id"`
	Notes     []types.Item `json:"notes"`
	Resources []types.Item `json:"resources"`
}

func toChapter(c *types.Chapter) *Chapter {
	if c == nil {
		return nil
	}
	result := Chapter{c.Id, c.Name, c.ParentId, []types.Item{}, []types.Item{}}

	for _, id := range c.Notes {
		result.Notes = append(result.Notes, storage.Store.Notes[id].Item)
	}

	return &result
}

func chapterHandler(w http.ResponseWriter, req *http.Request) {
	var result *Chapter
	data := []byte("null")

	if id := parseRequest(req, Paths["chapter"]); id != "" {
		s := storage.Store

		switch req.Method {
		case "GET":
			if s.Chapters[id] != nil {
				c := s.Chapters[id]
				result = toChapter(c)
			}
		case "POST":
			if id == "new" {
				c := &types.Chapter{}
				json.NewDecoder(req.Body).Decode(&c)

				if s.Books[c.ParentId] == nil {
					c = nil
				} else {
					c.Id = uuid.NewV4().String()
					s.Chapters[c.Id] = c
					s.Books[c.ParentId].AddChapter(c.Id)
					s.Save()
				}

				result = toChapter(c)
			}
		case "PATCH":
			if s.Chapters[id] != nil {
				c := s.Chapters[id]
				oldParent := c.ParentId
				json.NewDecoder(req.Body).Decode(&c)

				if s.Books[c.ParentId] == nil {
					c.ParentId = oldParent
				}

				if c.ParentId != oldParent {
					s.Books[c.ParentId].AddChapter(c.Id)
					s.Books[oldParent].RemoveChapter(c.Id)
				}
				s.Save()

				result = toChapter(c)
			}
		case "DELETE":
			// TODO: what about contained notes and resources
			if s.Chapters[id] != nil {
				c := s.Chapters[id]
				result = toChapter(c)

				delete(s.Chapters, c.Id)
				s.Books[c.ParentId].RemoveChapter(c.Id)

				s.Save()
			}
		}

		data, _ = json.Marshal(result)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
