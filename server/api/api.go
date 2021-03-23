package api

import (
	"log"
	"net/http"
	"strings"
)

var Paths = map[string]string{
	"store":   "/api/store/",
	"library": "/api/library/",
	"book":    "/api/book/",
	"chapter": "/api/chapter/",
	"note":    "/api/note/",
}

func Initialize(port string) {
	editorHandler := http.FileServer(http.Dir("./editor/dist"))
	http.Handle("/", editorHandler)

	http.HandleFunc(Paths["store"], storeHandler)
	http.HandleFunc(Paths["library"], libraryHandler)
	http.HandleFunc(Paths["book"], bookHandler)
	http.HandleFunc(Paths["chapter"], chapterHandler)
	http.HandleFunc(Paths["note"], noteHandler)

	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

////////////////////////////////////////////////////////////////////////////////

func parseRequest(req *http.Request, path string) string {
	elements := strings.Split(strings.TrimPrefix(req.URL.Path, path), "/")
	return elements[0]
}
