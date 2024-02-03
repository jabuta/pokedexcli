package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands((*config)(nil)) {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
