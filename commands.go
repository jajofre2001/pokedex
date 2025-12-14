package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jajofre2001/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

// Funcion que entrega un map de struct cliCommand disponibles
func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of the next 20 locations",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 locations",
			callback:    CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display a list of all the Pokemon located in the location specified",
			callback:    Explore,
		},
	}
}

// Comando que cierra el programa
func CommandExit(cfg *Config, arg []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

// Comando que entrega las funciones que se pueden hacer en el programa
func CommandHelp(cfg *Config, arg []string) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Println()
	fmt.Print("Usage:\n")
	fmt.Println()
	commands := GetCommands()

	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

// Comando que muestra en la consola 20 areas del mundo Pokemon
func CommandMap(cfg *Config, arg []string) error {
	resp_a, err := pokeapi.ListLocations()
	if err != nil {
		return err
	}

	cfg.Next = resp_a.Next
	cfg.Previous = resp_a.Previous

	for _, loc := range resp_a.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func CommandMapb(cfg *Config, arg []string) error {
	if cfg.Previous == nil {
		return errors.New("estas en la primera pagina")
	}

	resp_a, err := pokeapi.ListLocations()
	if err != nil {
		return err
	}

	cfg.Next = resp_a.Next
	cfg.Previous = resp_a.Previous

	for _, loc := range resp_a.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func Explore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location>")
	}

	location := args[0]

	specific_location_things, err := pokeapi.Specific_location_list(location)
	if err != nil {
		return err
	}

	pokemon_encounters := specific_location_things.PokemonEncounters

	for _, pokemon := range pokemon_encounters {
		fmt.Printf("Exploring %s...\n", location)
		fmt.Printf("Found Pokemon:\n")
		fmt.Printf("-%s\n", pokemon.Pokemon.Name)
	}
	return nil
}
