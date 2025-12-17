package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jajofre2001/pokedex/internal/pokeapi"
)

type Config struct {
	Next     *string
	Previous *string
	Pokedex  map[string]pokeapi.Pokemon
}

// Funcion que inicia el REPL y maneja los inputs del usuario
func StartRepl(cfg *Config) {
	user_input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		user_input.Scan()
		user_text := user_input.Text()
		clean_text := CleanInput(user_text)

		if len(clean_text) == 0 {
			continue
		}
		command_name := clean_text[0]
		args := clean_text[1:]

		command, exist := GetCommands()[command_name]
		if exist {
			err := command.callback(cfg, args)

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}

	}
}

// Funcion que recibe un str y regresa un slice str de palabras todas en minuscula
func CleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
