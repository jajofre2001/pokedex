package main

import (
	"reflect"
	"testing"
)

// Testeo de la funcion CleanInput
func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{ //Casos a testear
		"simple":           {input: "  hello  world  ", expected: []string{"hello", "world"}},
		"more things":      {input: "  Charmander Bulbasaur PIKACHU   ", expected: []string{"charmander", "bulbasaur", "pikachu"}},
		"una sola palabra": {input: "hola ", expected: []string{"hola"}},
		// add more cases here
	}

	for name, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("%s: expected: %v, got: %v", name, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if !reflect.DeepEqual(expectedWord, word) {
				t.Errorf("%s: expected: %v, got: %v", name, expectedWord, word)
			}

		}
	}
}
