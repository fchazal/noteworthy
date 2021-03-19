package server

import (
	"log"
	"net/http"

	"github.com/fchazal/noteworthy/server/books"
	"github.com/fchazal/noteworthy/server/chapters"
	"github.com/fchazal/noteworthy/server/libraries"
	"github.com/fchazal/noteworthy/server/notes"
	"github.com/fchazal/noteworthy/server/resources"
	"github.com/fchazal/noteworthy/server/storage"
)

func Start() {
	Initialize()

	editorHandler := http.FileServer(http.Dir("./editor/dist"))
	http.Handle("/", editorHandler)

	http.HandleFunc(libraries.Path, libraries.Handler)
	http.HandleFunc(books.Path, books.Handler)
	http.HandleFunc(chapters.Path, chapters.Handler)
	http.HandleFunc(notes.Path, notes.Handler)
	http.HandleFunc(resources.Path, resources.Handler)

	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Initialize() {
	memory := storage.Open("./noteworthy.json")
	memory.Save()
}
