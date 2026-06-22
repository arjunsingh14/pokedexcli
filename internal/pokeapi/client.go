package pokeapi

import (
	"net/http"
	"time"
	"github.com/arjunsingh14/pokedexcli/internal/cache"
)

type Client struct {
	httpClient http.Client
	cache *cache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(5 * time.Second),
	}
}