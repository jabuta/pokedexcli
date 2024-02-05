package main

import (
	"time"

	"github.com/jabuta/pokedexcli/internal/pokeAPI"
)

func main() {
	client := pokeAPI.NewClient(5*time.Second, 1*time.Minute)
	conf := &config{
		client:        client,
		caughtPokemon: make(map[string]pokeAPI.Pokemon),
	}
	startREPL(conf)
}
