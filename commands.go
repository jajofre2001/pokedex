package main

import (
	"errors"
	"fmt"
	"math/rand"
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
		"catch": {
			name:        "catch",
			description: "Try to capture a Pokemon",
			callback:    Catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the pokemos you had capture",
			callback:    Inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See the pokemons you have",
			callback:    Pokedex,
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

func Catch(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <Pokemon`s name>")
	}
	poke_name := args[0]

	pokemon, err := pokeapi.Request_pokemon(poke_name)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", poke_name)

	// Ver como implementar la posibilidad de captura
	chances := 1.0 / (1.0 + float64(pokemon.BaseExperience)/50.0)
	roll := rand.Float64()

	if roll < chances {
		fmt.Printf("%s was caught!\n", poke_name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
		cfg.Pokedex[poke_name] = pokemon

	} else {
		fmt.Printf("%s escaped\n", poke_name)
	}
	return nil
}

func Inspect(cfg *Config, arg []string) error {
	poke_name := arg[0]
	if pokemon, ok := cfg.Pokedex[poke_name]; !ok {
		return fmt.Errorf("you have not caught that pokemon")
	} else {

		fmt.Printf("Name: %s\n", poke_name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Printf("Types:\n")
		for _, tipo := range pokemon.Types {
			fmt.Printf("  -%s\n", tipo.Type.Name)
		}

	}
	return nil

}

func Pokedex(cfg *Config, arg []string) error {
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)

	}
	return nil
}
