package bookshelves

import (
	"encoding/json"
	"net/http"
)

const Path = "/api/libraries/"

func Handler(w http.ResponseWriter, req *http.Request) {
	var data []byte

	if req.Method == "GET" {

		data, _ = json.Marshal(nil)
	} else {
		element := New("name")
		json.NewDecoder(req.Body).Decode(element)

		switch req.Method {
		case "POST":
		case "PUT":
		case "DELETE":
		}

		data, _ = json.Marshal(&element)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
