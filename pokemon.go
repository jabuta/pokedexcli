package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a pokemon name or id")
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

func inspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a pokemon name")
	}
	pokemon, ok := conf.caughtPokemon[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}
	return nil
}

func pokedex(conf *config, args ...string) error {
	if len(conf.caughtPokemon) == 0 {
		return errors.New("you haven't caught any pokemon")
	}
	fmt.Println("Your Pokedex:")
	for k := range conf.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
