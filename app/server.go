package app

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	. "github.com/bonzofenix/package-calculator/processor"
)

// ## NewServer function
// returns an http.Handler
// it configures its own muxex and calls out to routes.go
func NewServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", RootHandler())
	return mux
}

func CalculateHandler(p IProcessor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsedOrder := 0
		parsedPackageSizes := []int{}

		if order := r.FormValue("order"); order != "" {
			parsedOrder, _ = strconv.Atoi(order)
		}

		// TODO: parsed package sizes, separte them by comma

		p.CalculatePacks(parsedPackageSizes, parsedOrder)

		tmp, err := template.New("Index").Parse(indexContent)

		if err != nil {
			log.Fatal(err)
		}
		tmp.Execute(w, nil)
	}
}

func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.New("Index").Parse(indexContent)
		if err != nil {
			log.Fatal(err)
		}

		tmp.Execute(w, nil)
	}
}
