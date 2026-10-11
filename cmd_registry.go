package main

import (
	"github.com/TheFredward/go_pokedex/internal/pokecache"
)

type config struct {
	commands        map[string]cliCommand
	NextURL         string
	PreviousURL     string
	cache           *pokecache.Cache
	capturedPokemon map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display 20 locations in Pokemon",
			callback:    getMapData,
		},
		"mapp": {
			name:        "mapp",
			description: "Displays the previous 20 locations in Pokemon",
			callback:    getPreviousMapData,
		},
		"explore": {
			name:        "explore",
			description: "Displays the name of all pokemon found in area",
			callback:    getLocationData,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to capture pokemon, requires user input",
			callback:    catchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "View data on specific pokemon. Usage inspect <pokemon-name>",
			callback:    inspectPokemon,
		},
		"list": {
			name:        "list",
			description: "Vew captured pokemon",
			callback:    listCaptured,
		},
	}
}
