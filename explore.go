package main

import (
	"errors"
	"fmt"
)

func explore(conf *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no location provided")
	}
	pokemonList, err := conf.client.ListPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	if pokemonList.Encounters[0].Pokemon.Name == "unown" {
		fmt.Println("No pokemon found")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, v := range pokemonList.Encounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}
	return nil
}
