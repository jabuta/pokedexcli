package pokeAPI

import (
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
