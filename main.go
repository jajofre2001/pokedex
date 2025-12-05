package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	user_input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		user_input.Scan()
		user_text := user_input.Text()
		clean_text := cleanInput(user_text)

		fmt.Printf("Tu comando fue: %v\n", clean_text[0])
	}

}

// Funcion que recibe un str y regresa un slice str de palabras todas en minuscula
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
