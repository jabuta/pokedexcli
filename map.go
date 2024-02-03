package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jabuta/pokedexcli/internal/pokeAPI"
)

func commandMap(conf *config) func() error {
	return func() error {

		var endpointURL string
		if conf.next == "" {
			endpointURL = pokeAPI.GetEndPoints()["map"]
		} else {
			endpointURL = conf.next
		}

		responseBody, err := pokeAPI.QueryAPI(endpointURL)
		if err != nil {
			return err
		}

		locationList := LocationResponse{}
		err = json.Unmarshal(responseBody, &locationList)
		if err != nil {
			return err
		}
		conf.next = locationList.Next
		conf.prev = locationList.Previous
		printLocationResults(locationList)
		return nil
	}
}

func commandMapB(conf *config) func() error {
	return func() error {

		var endpointURL string
		if conf.prev == "" || conf.prev == "null" {
			return errors.New("error: you're on the first page")
		} else {
			endpointURL = conf.prev
		}

		responseBody, err := pokeAPI.QueryAPI(endpointURL)
		if err != nil {
			return err
		}

		locationList := LocationResponse{}
		err = json.Unmarshal(responseBody, &locationList)
		if err != nil {
			return err
		}
		conf.next = locationList.Next
		conf.prev = locationList.Previous
		printLocationResults(locationList)
		return nil
	}
}

func printLocationResults(locationlist LocationResponse) {
	for _, loc := range locationlist.Results {
		fmt.Println(loc.Name)
	}
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
