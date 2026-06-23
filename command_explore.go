package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Requires an area id or name. Use map command to find an area")
		return nil
	}
	id := args[0]
	locationAreaDetail, err := cfg.client.GetLocationAreaDetail(id)

	if err != nil {
		return err
	}
	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")
	for _, e := range locationAreaDetail.PokemonEncounters {
		fmt.Printf("- %s\n", e.Pokemon.Name)
	}

	return nil
}
