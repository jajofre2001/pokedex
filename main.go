package main

import (
	"math/rand"
	"time"

	"github.com/jajofre2001/pokedex/internal/pokeapi"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := &Config{
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	StartRepl(cfg)
}
