package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/TheFredward/go_pokedex/internal/pokecache"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		cache:    pokecache.NewCache(5 * time.Second),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		userInput := cleanInput(scanner.Text())
		cmd, exists := cfg.commands[userInput[0]]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
