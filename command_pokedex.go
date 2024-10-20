package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) < 1 {
		return errors.New("you haven't caught any pokemon yet")

	}

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("%v - %v\n", pokemon.ID, pokemon.Name)
	}

	return nil
}
