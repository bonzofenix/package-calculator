package app

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	. "github.com/bonzofenix/package-calculator/processor"
)

// ## NewServer function
// returns an http.Handler
// it configures its own muxex and calls out to routes.go
func NewServer() http.Handler {
	mux := http.NewServeMux()
	processor := NewProcessor()
	mux.HandleFunc("/", RootHandler())
	mux.HandleFunc("/calculate", CalculateHandler(processor))
	return mux
}

func CalculateHandler(p IProcessor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var packSizes, order string
		parsedOrder := 0
		parsedPackageSizes := []int{}

		if order = r.FormValue("order"); order != "" {
			parsedOrder, _ = strconv.Atoi(order)
		}

		// Parsed package sizes, separte them by comma
		if packSizes = r.FormValue("packSizes"); packSizes != "" {
			packs := strings.Split(packSizes, ",")
			parsedPackageSizes = make([]int, len(packs))
			for i, pack := range packs {
				parsedPackageSizes[i], _ = strconv.Atoi(pack)
			}
		}

		if len(parsedPackageSizes) == 0 || parsedOrder == 0 {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		result := p.CalculatePacks(parsedPackageSizes, parsedOrder)
		tmp, err := template.New("Index").Parse(indexContent)

		if err != nil {
			log.Fatal(err)
		}
		tmp.Execute(w, struct {
			Order     string
			PackSizes string
			Result    map[int]int
		}{
			order,
			packSizes,
			result,
		})
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
