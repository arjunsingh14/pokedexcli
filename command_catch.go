package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args... string) error {
	if len(args) == 0 {
		fmt.Println("Enter a valid Pokemon name")
		return nil
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.client.GetPokemon(name)
	if err != nil {
		return err
	}
	const threshold = 50
	roll := rand.Intn(pokemon.BaseExperience)
	if roll < threshold {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
