package main

import (
	"fmt"
	"os"
)

func commandHelp(conf *config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(conf *config) error {
	os.Exit(0)
	return nil
}
