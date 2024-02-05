package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(conf *config, args []string) error {
	if len(args) != 1 {
		return errors.New("provide a location name or id")
	}
	pokemon, err := conf.client.FetchPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := (608.0-float64(pokemon.BaseExperience))/608.0 + rand.Float64()
	if chance < 1 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	conf.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
