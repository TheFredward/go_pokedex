package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var val []string
	lowerCase := strings.ToLower(text)
	cleanedString := strings.TrimSpace(lowerCase)
	val = strings.Split(cleanedString, " ")

	return val
}
