package main

import (
	"errors"
	"fmt"

	"github.com/jabuta/pokedexcli/internal/pokeAPI"
)

func commandMap(conf *config) error {
	if conf.next == nil && conf.prev != nil {
		fmt.Println("End of List, try mapb")
		return errors.New("end of list")
	}

	locationList, err := conf.client.ListLocations(conf.next)
	if err != nil {
		return err
	}

	conf.next = locationList.Next
	conf.prev = locationList.Previous
	printLocationResults(locationList)
	return nil
}

func commandMapB(conf *config) error {
	if conf.prev == nil {
		fmt.Println("Start of List, try map")
		return errors.New("start of list")
	}
	locationList, err := conf.client.ListLocations(conf.prev)
	if err != nil {
		return err
	}

	conf.next = locationList.Next
	conf.prev = locationList.Previous
	printLocationResults(locationList)
	return nil
}

func printLocationResults(locationlist pokeAPI.LocationResponse) {
	for _, loc := range locationlist.Results {
		fmt.Println(loc.Name)
	}
}
