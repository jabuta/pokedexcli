package main

import (
	"errors"
	"fmt"
)

func explore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a location name or id")
	}
	location, err := conf.client.ListPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	if location.Encounters[0].Pokemon.Name == "unown" {
		fmt.Println("No pokemon found")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, v := range location.Encounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}
	return nil
}
