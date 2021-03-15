package server

import "log"
import "net/http"

import Api "github.com/fchazal/noteworthy/server/api"

func Start () {
  staticHandler := http.FileServer(http.Dir("./editor/dist"))
  http.Handle("/", staticHandler)

	http.HandleFunc("/api/", Api.Handler)
	http.HandleFunc("/api/libraries/", Api.LibrariesHandler)
	http.HandleFunc("/api/books/", Api.BooksHandler)
	http.HandleFunc("/api/chapters/", Api.ChaptersHandler)
	http.HandleFunc("/api/notes/", Api.NotesHandler)
	http.HandleFunc("/api/resources/", Api.ResourcesHandler)

  log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}