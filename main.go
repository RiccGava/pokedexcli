package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the cli interface",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to Pokedex")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("exiting program")
	os.Exit(0)
	return nil
}

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if cmd, exists := commands[input]; exists {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command", input)
		}

	}
}
