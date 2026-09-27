package main

import (
	"fmt"
	"os"
	"strings"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	cmd := getCommands()
	for i := range cmd {
		fmt.Printf("%s: %s\n", cmd[i].name, cmd[i].description)
	}
	return nil
}

func cleanInput(text string) []string {
	var val []string
	lowerCase := strings.ToLower(text)
	cleanedString := strings.TrimSpace(lowerCase)
	val = strings.Split(cleanedString, " ")

	return val
}
