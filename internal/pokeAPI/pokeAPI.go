package pokeAPI

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/jabuta/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, cacheTtl time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheTtl),
	}
}

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageURL *string) (LocationResponse, error) {
	endPoint := baseURL + "/location-area"
	if pageURL != nil {
		endPoint = *pageURL
	}

	var responseBody []byte
	responseBody, ok := c.cache.Get(endPoint)
	if !ok {

		req, err := http.NewRequest("GET", endPoint, nil)
		if err != nil {
			return LocationResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationResponse{}, err
		}

		responseBody, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return LocationResponse{}, err
		}
		if res.StatusCode > 299 {
			return LocationResponse{}, errors.New(res.Status)
		}

		c.cache.Add(endPoint, responseBody)
	}

	locationList := LocationResponse{}
	err := json.Unmarshal(responseBody, &locationList)
	if err != nil {
		return LocationResponse{}, err
	}
	return locationList, nil
}

type LocationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
