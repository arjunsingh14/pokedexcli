package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	val, exists := c.cache.Get(url)
	if exists {
		var pokemon Pokemon
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var poke Pokemon
	if err := json.Unmarshal(body, &poke); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return poke, nil
}
