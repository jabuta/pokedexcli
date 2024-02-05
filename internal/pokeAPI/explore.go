package pokeAPI

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location string) (Location, error) {
	if location == "" {
		return Location{}, errors.New("no location")
	}
	endPoint := baseURL + "/location-area/" + location

	var responseBody []byte
	responseBody, ok := c.cache.Get(endPoint)
	if !ok {

		req, err := http.NewRequest("GET", endPoint, nil)
		if err != nil {
			return Location{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}

		responseBody, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return Location{}, err
		}
		if res.StatusCode > 299 {
			return Location{}, errors.New(res.Status)
		}

		c.cache.Add(endPoint, responseBody)
	}

	pokemonList := Location{}
	err := json.Unmarshal(responseBody, &pokemonList)
	if err != nil {
		return Location{}, err
	}
	return pokemonList, nil

}

type Location struct {
	Name       string `json:"name"` // Assuming "name" is the field for location name
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
