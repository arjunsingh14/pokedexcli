package main

import (
	"fmt"
	"github.com/arjunsingh14/pokedexcli/internals"
)


func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	locationArea, err := internals.GetLocationAreas(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = locationArea.Next
	cfg.previous = locationArea.Previous

	for _, result := range locationArea.Results{
		fmt.Println(result.Name)
	}

	return nil
}