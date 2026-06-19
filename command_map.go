package main

import (
	"fmt"
	"github.com/arjunsingh14/pokedexcli/internals"
)



func commandMap(cfg *config) error {
	locationArea, err := internals.GetLocationAreas(cfg.next)
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