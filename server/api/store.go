package api

import (
	"encoding/json"
	"net/http"

	"github.com/fchazal/noteworthy/server/storage"
	"github.com/fchazal/noteworthy/server/types"
)

////////////////////////////////////////////////////////////////////////////////

type Store struct {
	Libraries []types.Item `json:"libraries"`
	Books     []types.Item `json:"books"`
}

func storeHandler(w http.ResponseWriter, req *http.Request) {
	data := []byte("null")
	result := Store{[]types.Item{}, []types.Item{}}

	if id := parseRequest(req, Paths["store"]); req.Method == "GET" && id == "" {
		s := storage.Store

		for _, id := range s.BookStore.Libraries {
			e := s.Libraries[id]
			result.Libraries = append(result.Libraries, e.Item)
		}

		for _, id := range s.BookStore.Books {
			e := s.Books[id]
			result.Books = append(result.Books, e.Item)
		}

		data, _ = json.Marshal(result)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
