package chapters

import "net/http"

const Path = "/api/chapters/"

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

	} else {
		switch req.Method {
		case "POST":
		case "PUT":
		case "DELETE":
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(nil)
}
