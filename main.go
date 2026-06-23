package main

import (
	"github.com/arjunsingh14/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		client: client,
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
