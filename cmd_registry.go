package main

type config struct {
	commands    map[string]cliCommand
	NextURL     string
	PreviousURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}
