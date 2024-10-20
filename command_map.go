package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println(("Locations areas:"))
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return nil
}
