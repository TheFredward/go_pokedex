package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var val []string
	cleanedString := strings.TrimSpace(text)
	val = strings.Split(cleanedString, " ")

	return val
}
