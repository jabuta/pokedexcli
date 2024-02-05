package pokeAPI

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location string) (PokemonResponse, error) {
	if location == "" {
		return PokemonResponse{}, errors.New("no location")
	}
	endPoint := baseURL + "/location-area/" + location

	var responseBody []byte
	responseBody, ok := c.cache.Get(endPoint)
	if !ok {

		req, err := http.NewRequest("GET", endPoint, nil)
		if err != nil {
			return PokemonResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonResponse{}, err
		}

		responseBody, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return PokemonResponse{}, err
		}
		if res.StatusCode > 299 {
			return PokemonResponse{}, errors.New(res.Status)
		}

		c.cache.Add(endPoint, responseBody)
	}

	pokemonList := PokemonResponse{}
	err := json.Unmarshal(responseBody, &pokemonList)
	if err != nil {
		return PokemonResponse{}, err
	}
	return pokemonList, nil

}

type PokemonResponse struct {
	LocationName string `json:"name"` // Assuming "name" is the field for location name
	Encounters   []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
