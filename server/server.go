package server

import (
	"log"
	"net/http"

	"github.com/fchazal/noteworthy/server/books"
	"github.com/fchazal/noteworthy/server/bookshelves"
	"github.com/fchazal/noteworthy/server/storage"
)

func Start() {
	Initialize()
	return

	editorHandler := http.FileServer(http.Dir("./editor/dist"))
	http.Handle("/", editorHandler)

	http.HandleFunc(bookshelves.Path, bookshelves.Handler)
	http.HandleFunc(books.Path, books.Handler)
	//	http.HandleFunc(chapters.Path, chapters.Handler)
	//	http.HandleFunc(notes.Path, notes.Handler)
	//	http.HandleFunc(resources.Path, resources.Handler)

	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Initialize() {
	storage.Open("./noteworthy.json")

	shelf := bookshelves.New("Étagère")
	storage.Library.AddBookshelf(shelf)

	b1 := books.New("Non")
	shelf.AddBook(b1)

	storage.Save()
}
