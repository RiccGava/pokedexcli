package main

import (
	"errors"
	"fmt"
)

func callbackMapb(cfg *config) error {

	if cfg.previousLocationAreaURL == nil {

		return errors.New("you can't go back from here")
	}
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return nil
}
