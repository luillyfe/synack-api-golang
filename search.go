package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	log.Printf("Calling %s with query parameters, %v", r.URL.Path[1:], r.URL.Query())
	q := r.URL.Query()
	query := q.Get("query")
	engine := q.Get("engine")

	var items []byte

	switch engine {
		case "GOOGLE":
			items = searchOnGoogle(query)
		case "BING":
			items = searchOnBing(query)
		case "BOTH":
			items = searchOnBoth(query)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(items)
}