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
}

func Initialize(port string) {
	editorHandler := http.FileServer(http.Dir("./editor/dist"))
	http.Handle("/", editorHandler)

	http.HandleFunc(Paths["store"], storeHandler)
	http.HandleFunc(Paths["library"], libraryHandler)
	http.HandleFunc(Paths["book"], bookHandler)

	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

////////////////////////////////////////////////////////////////////////////////

func parseRequest(req *http.Request, path string) string {
	elements := strings.Split(strings.TrimPrefix(req.URL.Path, path), "/")
	return elements[0]
}
