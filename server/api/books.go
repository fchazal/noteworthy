package api

import "net/http"

type Book struct {
	Id		string
	Name	string
}

var Books = []Books{}

func (b *Book) update() {
	
}

func (b *Book) move() {

}

func (b *Book) delete() {

}


func BooksHandler(w http.ResponseWriter, req *http.Request) {
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