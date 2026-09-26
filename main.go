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
		var splitInput []string
		splitInput = cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", splitInput[0])
	}
}
