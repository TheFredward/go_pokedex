package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type queriedJsonData struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var baseURL string = "https://pokeapi.co/api/v2/"

func getResults(cfg *config) {

}

func getMapData(cfg *config) error {
	// if cfg.Offset is unassigned or empty set it as zero
	// each time after add or subtrach by 20 as long as the result is greather than zero
	var resultsJson queriedJsonData
	var err error
	var storedData []byte
	var res *http.Response
	if cfg.NextURL == "" {
		cfg.NextURL = baseURL + "location-area/?limit=20&offset=0"
	}

	storedData, ok := cfg.cache.Get(cfg.NextURL)
	if ok {
		fmt.Print("CACHE HIT!!!")
		err = json.Unmarshal(storedData, &resultsJson)

	} else {
		var clientCall = &http.Client{Timeout: 10 * time.Second}
		fullURL := cfg.NextURL
		res, err = clientCall.Get(fullURL)
		if err != nil {
			fmt.Print(err)
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			fmt.Printf("Response failed with:\nstatus code: %d\nbody: %s", res.StatusCode, res.Body)
		}
		storedData, err = io.ReadAll(res.Body)
		cfg.cache.Add(cfg.NextURL, storedData)
	}
	if err != nil {
		fmt.Printf("Unable to decode response body: %+v", err)
	}
	err = json.Unmarshal(storedData, &resultsJson)
	if err != nil {
		fmt.Printf("Unable to decode response body: %+v", err)
	}
	for _, mapLocation := range resultsJson.Results {
		fmt.Printf("%s\n", mapLocation.Name)
	}
	cfg.NextURL = resultsJson.Next
	cfg.PreviousURL = resultsJson.Previous

	return nil
}

func getPreviousMapData(cfg *config) error {
	// if cfg.Offset is unassigned or empty set it as zero
	// each time after add or subtrach by 20 as long as the result is greather than zero
	var resultsJson queriedJsonData
	var res *http.Response
	var err error
	var storedData []byte
	if cfg.PreviousURL == "" {
		fmt.Print("You are on the first page!\n")
		return nil
	}
	storedData, ok := cfg.cache.Get(cfg.PreviousURL)
	if ok {
		fmt.Print("CACHE HIT!!!")
		err = json.Unmarshal(storedData, &resultsJson)
	} else {
		var clientCall = &http.Client{Timeout: 10 * time.Second}
		fullURL := cfg.PreviousURL
		res, err = clientCall.Get(fullURL)
		if err != nil {
			fmt.Print(err)
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			fmt.Printf("Response failed with:\nstatus code: %d\nbody: %s", res.StatusCode, res.Body)
		}
		storedData, err = io.ReadAll(res.Body)
		cfg.cache.Add(cfg.PreviousURL, storedData)
	}
	if err != nil {
		fmt.Printf("Unable to decode response body: %+v", err)
	}
	err = json.Unmarshal(storedData, &resultsJson)
	if err != nil {
		fmt.Printf("Unable to decode response body: %+v", err)
	}
	for _, mapLocation := range resultsJson.Results {
		fmt.Printf("%s\n", mapLocation.Name)
	}
	cfg.PreviousURL = resultsJson.Previous
	cfg.NextURL = resultsJson.Next

	return nil
}
