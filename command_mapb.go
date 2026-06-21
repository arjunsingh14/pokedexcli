package main

import (
	"fmt"
)


func commandMapb(cfg *config) error {
	if *cfg.previousLocationsUrl == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	locationArea, err := cfg.client.GetLocationAreas(cfg.previousLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationArea.Next
	cfg.previousLocationsUrl = locationArea.Previous

	for _, result := range locationArea.Results{
		fmt.Println(result.Name)
	}

	return nil
}