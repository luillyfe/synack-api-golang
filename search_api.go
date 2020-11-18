package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func searchOnBing(query string) (responseJson []byte) {
	response := BingResponse{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", viperEnvVariable("URI_BING") + query, nil)
	req.Header.Add("Ocp-Apim-Subscription-Key", viperEnvVariable("BING_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	responseJson, _ = json.Marshal(response.WebPages)
	return
}

func searchOnGoogle(query string) (responseJson []byte) {
	response := GoogleResponse{}

	resp, err := http.Get(viperEnvVariable("URI_GOOGLE") + query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	responseJson, _ = json.Marshal(response)
	return
}

func searchOnBoth(query string) []byte {
	ch := make(chan []byte)
	defer close(ch)

	go func(ch chan []byte) {
		ch <- searchOnBing(query)
	}(ch)

	go func(ch chan []byte) {
		ch <- searchOnGoogle(query)
	}(ch)

	return append(<-ch, <-ch...)
}

type GoogleResponse struct {
	Items []GoogleItems `json:"items"`
}

type BingResponse struct {
	WebPages BingValue
}

type GoogleItems struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

type BingValue struct {
	Items []BingItems `json:"value"`
}

type BingItems struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Snippet string `json:"snippet"`
}
