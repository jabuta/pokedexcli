package main

import (
	"time"

	"github.com/jabuta/pokedexcli/internal/pokeAPI"
)

func main() {
	client := pokeAPI.NewClient(500 * time.Millisecond)
	conf := &config{
		client: client,
	}
	startREPL(conf)
}
