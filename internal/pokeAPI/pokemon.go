package pokeAPI

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(pokemonName string) (Pokemon, error) {
	if pokemonName == "" {
		return Pokemon{}, errors.New("no pokemon name")
	}
	endPoint := baseURL + "/pokemon/" + pokemonName

	var responseBody []byte
	responseBody, ok := c.cache.Get(endPoint)
	if !ok {

		req, err := http.NewRequest("GET", endPoint, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		responseBody, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return Pokemon{}, err
		}
		if res.StatusCode > 299 {
			return Pokemon{}, errors.New(res.Status)
		}

		c.cache.Add(endPoint, responseBody)
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(responseBody, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

type Pokemon struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []PokemonAbility   `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []VersionGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	Species                NamedAPIResource   `json:"species"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
	PastTypes              []PokemonTypePast  `json:"past_types"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type PokemonHeldItem struct {
	Item           NamedAPIResource         `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type PokemonMove struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
}

type PokemonStat struct {
	Stat     NamedAPIResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []PokemonType    `json:"types"`
}
