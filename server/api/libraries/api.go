package libraries

import (
	"encoding/json"
	"net/http"

	"github.com/fchazal/noteworthy/server/utils"
	uuid "github.com/satori/go.uuid"
)

const Path = "/api/libraries/"

func Handler(w http.ResponseWriter, req *http.Request) {
	var data []byte

	if req.Method == "GET" {
		data, _ = json.Marshal(Libraries)
	} else {
		l := Library{}
		json.NewDecoder(req.Body).Decode(&l)

		l.Id = uuid.NewV4().String()
		l.Path = utils.ToPath(l.Name)

		switch req.Method {
		case "POST":
			addLibrary(&l)

		case "PUT":
		case "DELETE":

		}

		data, _ = json.Marshal(&l)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
