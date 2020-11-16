package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleSearch)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	log.Printf("Calling %s with query parameters, %v", r.URL.Path[1:], r.URL.Query())
	query := r.URL.Query()
	response := Response{}

	resp, err := http.Get(viperEnvVariable("URI_GOOGLE") + query.Get("query"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	responseJson, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

type Response struct {
	Items []Items `json:"items"`
}

type Items struct {
	Title string `json:"title"`
	Link string `json:"link"`
	Snippet string `json:"snippet"`
}