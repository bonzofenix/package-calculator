package app

import (
	"html/template"
	"log"
	"net/http"
)

// ## NewServer function
// returns an http.Handler
// it configures its own muxex and calls out to routes.go
func NewServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", RootHandler)
	return mux
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.New("Index").Parse(indexContent)
	if err != nil {
		log.Fatal(err)
	}

	tmp.Execute(w, nil)
}
