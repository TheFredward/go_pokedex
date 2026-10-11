package main

type PokemonDict struct {
	Pokemon PokemonName `json:"pokemon"`
}
type PokemonName struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Types          []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		}
	}
	Abilities []struct {
		Is_hidden bool `json:"is_hidden"`
		Slot      int  `json:"slot"`
		Ability   struct {
			Name string `json:"name"`
		}
	}
}
