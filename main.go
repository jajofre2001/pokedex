package main

import (
	"math/rand"
	"time"

	"github.com/jajofre2001/pokedex/internal/pokeapi"
)

// Funcion que da inicia el programa
func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := &Config{
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	StartRepl(cfg)
}
