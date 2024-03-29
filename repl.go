package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jabuta/pokedexcli/internal/pokeAPI"
)

func startREPL(conf *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		if cmd, ok := getCommands()[command]; ok {
			if err := cmd.callback(conf, words[1:]...); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command does not exist")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	client        pokeAPI.Client
	caughtPokemon map[string]pokeAPI.Pokemon
	next          *string
	prev          *string
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Retrieve the next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Retrieve the next locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Find the pokemon in a given area",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "View info in a pokemon you caught",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View the pokemon you've caught",
			callback:    pokedex,
		},
	}
}
