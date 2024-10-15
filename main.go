package main

import (
	"fmt"
	"log"

	"github.com/riccgava/pokedexcli/internal/pokeapi"
)

func main() {

	//startRepl()

	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationArea()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
