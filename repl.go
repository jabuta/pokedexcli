package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {

	scanner := bufio.NewScanner(os.Stdin)
	conf := &config{}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		if cmd, ok := getCommands(conf)[command]; ok {
			if err := cmd.callback(); err != nil {
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
	callback    func() error
}

type config struct {
	next string
	prev string
}

func getCommands(conf *config) map[string]cliCommand {

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
			callback:    commandMap(conf),
		},
		"mapb": {
			name:        "map",
			description: "Retrieve the next locations",
			callback:    commandMapB(conf),
		},
	}
}
