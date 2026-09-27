package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "deleteafter reading this block",
			expected: []string{"deleteafter", "reading", "this", "block"},
		},
		{
			input:    "DELEtEafter readingthis     ",
			expected: []string{"DELEtEafter", "readingthis"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Mismatch length %v vs %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				// if they don't match, use t.Errorf to print an error message
				t.Errorf("Mismatch word! %v vs %v", word, expectedWord)
			}
		}
	}
}
