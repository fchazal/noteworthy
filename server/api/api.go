package api

import "fmt"
import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/api/"):]

	fmt.Fprintf(w, "<h1>%s</h1>", title)
}