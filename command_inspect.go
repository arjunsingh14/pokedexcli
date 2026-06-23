package main

import "fmt"

func commandInspect(cfg *config, args... string) error {
	if len(args) == 0 {
		fmt.Println("Please enter a valid Pokemon name")
		return nil
	}
	name := args[0]
	pokedex := cfg.pokedex
	pokemon, exists := pokedex[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}