package api

import (
	"log"
	"net/http"

	"github.com/fchazal/noteworthy/server/api/books"
	"github.com/fchazal/noteworthy/server/api/chapters"
	"github.com/fchazal/noteworthy/server/api/libraries"
	"github.com/fchazal/noteworthy/server/api/notes"
	"github.com/fchazal/noteworthy/server/api/resources"
	"github.com/fchazal/noteworthy/server/storage"
)

func Initialize(port string) {
	// TODO: Ugly hack to achieve this without circular dependencies
	save := storage.Save

	editorHandler := http.FileServer(http.Dir("./editor/dist"))
	http.Handle("/", editorHandler)

	http.HandleFunc(libraries.Path, libraries.Handler)
	libraries.Libraries = &(storage.Library.Libraries)
	libraries.Save = &(save)

	http.HandleFunc(books.Path, books.Handler)
	books.Books = &(storage.Library.Books)
	books.Save = &(save)

	http.HandleFunc(chapters.Path, chapters.Handler)
	chapters.Chapters = &(storage.Library.Chapters)
	chapters.Save = &(save)

	http.HandleFunc(notes.Path, notes.Handler)
	notes.Notes = &(storage.Library.Notes)
	notes.Save = &(save)

	http.HandleFunc(resources.Path, resources.Handler)
	resources.Resources = &(storage.Library.Resources)
	resources.Save = &(save)

	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
