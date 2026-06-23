package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationAreaDetail(id string) (locationAreaDetail, error ){
	url := baseURL + "/location-area/" + id
	val, exists := c.cache.Get(url)
	if exists {
		var locationAreaDetail locationAreaDetail
		err := json.Unmarshal(val, &locationAreaDetail)
		if err != nil {
			return locationAreaDetail, err
		}
		return locationAreaDetail, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaDetail{}, err
	}

  res, err := c.httpClient.Do(req)
  if err != nil {
      return locationAreaDetail{}, err
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
      return locationAreaDetail{}, err
  }

  var payload locationAreaDetail
  if err := json.Unmarshal(body, &payload); err != nil {
      return locationAreaDetail{}, err
  }

  c.cache.Add(url, body)
  return payload, nil
} 
