package api

import "log"
import "net/http"
import "encoding/json"

import uuid "github.com/satori/go.uuid"

type Library struct {
	Id		string
	Name	string
	Books	[]Book
}

var Libraries = []Library{}


func (l *Library) initialize() {
	l.Id = uuid.NewV4().String()
}

func (l *Library) AddBook(book *Book) {

}

func newLibrary(name string) Library {
	l := Library{
		uuid.NewV4().String(),
		name,
		[]Book{},
	}
	return l
}

func LibrariesHandler(w http.ResponseWriter, req *http.Request) {
	var data	[]byte

	if req.Method == "GET" {
		log.Println("GET")

		elements := []Library{
			newLibrary("riri"),
			newLibrary("fifi"),
			newLibrary("loulou"),
		}


		data, _ = json.Marshal(&elements)
	} else {
		var element	Library
		element.initialize()
		json.NewDecoder(req.Body).Decode(&element)
		
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