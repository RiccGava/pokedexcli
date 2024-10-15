package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to pokedex")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %v: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
