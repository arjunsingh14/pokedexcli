package main

import (
	"fmt"
)



func commandMap(cfg *config, _ ...string) error {
	locationArea, err := cfg.client.GetLocationAreas(cfg.nextLocationsUrl)
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