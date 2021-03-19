package books

import (
	"encoding/json"
	"net/http"
	"strings"
)

const Path = "/api/books/"

func Handler(w http.ResponseWriter, req *http.Request) {
	result := []byte("ERROR")
	elements := strings.Split(strings.TrimPrefix(req.URL.Path, Path), "/")

	if req.Method == "GET" {
		if len(elements) == 1 {
			if (*Items)[elements[0]] != nil {
				result, _ = json.Marshal((*Items)[elements[0]])
			}
		}
	} else {
		switch req.Method {
		case "POST":
			if len(elements) == 1 && elements[0] == "new" {

			}
		case "PUT":
		case "DELETE":
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
