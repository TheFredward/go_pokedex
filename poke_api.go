package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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

type locationJsonData struct {
	Pokemon []PokemonDict `json:"pokemon_encounters"`
}

var baseURL string = "https://pokeapi.co/api/v2/"

func getMapData(cfg *config, userInput string) error {
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

func getPreviousMapData(cfg *config, userInput string) error {
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

func getLocationData(cfg *config, location string) error {
	var resultsJson locationJsonData
	var err error
	var res *http.Response
	var storedData []byte
	var clientCall = &http.Client{Timeout: 10 * time.Second}
	fullURL := baseURL + "location-area/" + location
	res, err = clientCall.Get(fullURL)
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with:\nstatus code: %d\nbody: %s", res.StatusCode, res.Body)
	}
	storedData, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to decode response body: %+v", err)
	}
	err = json.Unmarshal(storedData, &resultsJson)
	if err != nil {
		fmt.Printf("Unable to decode unmarshall: %+v", err)
	}
	for _, innerObj := range resultsJson.Pokemon {
		fmt.Printf("%s\n", innerObj.Pokemon.Name)
	}

	return nil

}

func catchPokemon(cfg *config, pokeName string) error {
	var clientCall = &http.Client{Timeout: 10 * time.Second}
	var res *http.Response
	var err error
	var storedData []byte
	var PokemonData Pokemon
	randomChance := rand.Intn(300)
	userPokemon := cfg.capturedPokemon
	fullUrl := baseURL + "pokemon/" + pokeName
	res, err = clientCall.Get(fullUrl)
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response Failed with Status Code: %d\nBody: %s", res.StatusCode, res.Body)
	} else {
		storedData, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Unable to read body response: %v", err)
		}
		fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
		err = json.Unmarshal(storedData, &PokemonData)
		if err != nil {
			fmt.Printf("Unable to unmarshall: %v", err)
		}
		if randomChance >= PokemonData.BaseExperience {
			fmt.Printf("%s was caught!\n", PokemonData.Name)
			userPokemon[PokemonData.Name] = PokemonData
		} else {
			fmt.Printf("%s escaped!\n", PokemonData.Name)
		}

	}
	return nil
}
func listCaptured(cfg *config, userString string) error {
	if len(cfg.capturedPokemon) == 0 {
		fmt.Print("No pokemon captured!\n")
	} else {
		for monster := range cfg.capturedPokemon {
			fmt.Printf("%s ", monster)
		}
	}
	return nil
}

func inspectPokemon(cfg *config, name string) error {
	if len(name) == 0 {
		fmt.Printf("Pokemon name is needed!: %s", cfg.commands["inspection"].description)
		return nil
	}
	var clientCall = &http.Client{Timeout: 10 * time.Second}
	var response *http.Response
	var err error
	var pokemon Pokemon
	fullUrl := baseURL + "pokemon/" + name
	response, err = clientCall.Get(fullUrl)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
	}
	if response.StatusCode > 299 {
		fmt.Printf("Response Failed with Status Code: %d\nBody: %s", response.StatusCode, string(body))
		return nil
	} else {
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			fmt.Printf("Unable to unmarshall: %v", err)
		}
		var pType string
		var pAbility string
		for _, val := range pokemon.Types {
			pType = val.Type.Name
		}
		for _, val := range pokemon.Abilities {
			if !val.Is_hidden {
				pAbility = val.Ability.Name
				break
			}
		}
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nBase Experience: %d\nType: %s\nNon-Hidden Ability: %s\n", pokemon.Name, pokemon.Height, pokemon.Weight, pokemon.BaseExperience, pType, pAbility)
	}
	return nil
}
