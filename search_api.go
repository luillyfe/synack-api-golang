package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func searchOnBing(query string) (responseJson []byte) {
	response := bingResponse{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", viperEnvVariable("URI_BING")+query, nil)
	if err != nil {
		log.Fatal(err)
	}
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

	responseJson, _ = json.Marshal(response.WebPages.Value)
	bingItems := make([]bingItems, 0)
	err = json.Unmarshal(responseJson, &bingItems)
	if err != nil {
		log.Fatal(err)
	}

	out := make([]googleItems, 0, len(bingItems))
	for _, o := range bingItems {
		out = append(out, googleItems{Title: o.Name, Link: o.Url, Snippet: o.Snippet})
	}
	responseJson, _ = json.Marshal(googleResponse{Items: out})
	return
}

func searchOnGoogle(query string) (responseJson []byte) {
	response := googleResponse{}

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
	ch := make(chan []googleItems)
	defer close(ch)

	go func(ch chan []googleItems) {
		var bing googleResponse

		err := json.Unmarshal(searchOnBing(query), &bing)
		if err != nil {
			log.Fatal(err)
		}

		ch <- bing.Items
	}(ch)

	go func(ch chan []googleItems) {
		var google googleResponse

		err := json.Unmarshal(searchOnGoogle(query), &google)
		if err != nil {
			log.Fatal(err)
		}

		ch <- google.Items
	}(ch)

	responseJSON, _ := json.Marshal(&response{Bing: <-ch, Google: <-ch})
	return responseJSON
}

type response struct {
	Google []googleItems `json:"google"`
	Bing []googleItems `json:"bing"`
}

type googleResponse struct {
	Items []googleItems `json:"items"`
}

type bingResponse struct {
	WebPages bingValue `json:"webPages"`
}

type googleItems struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

type bingValue struct {
	Value []bingItems `json:"value"`
}

type bingItems struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Snippet string `json:"snippet"`
}
