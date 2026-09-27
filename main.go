package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		userInput := cleanInput(scanner.Text())
		cmd, exists := getCommands()[userInput[0]]
		if exists {
			err := cmd.callback()
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
