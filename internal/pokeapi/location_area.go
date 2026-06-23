package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationArea struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []result `json:"results"`
}

func (c *Client) GetLocationAreas(pageURL *string) (locationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	val, exists := c.cache.Get(url)
	if exists {
		var locationAreas locationArea
		err := json.Unmarshal(val, &locationAreas)
		if err != nil {
			return locationAreas, err
		}
		return locationAreas, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}

	res, err := c.httpClient.Do(req)
  if err != nil {
      return locationArea{}, err
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
      return locationArea{}, err
  }

  var locationAreas locationArea
  if err := json.Unmarshal(body, &locationAreas); err != nil {
      return locationArea{}, err
  }

  c.cache.Add(url, body)
  return locationAreas, nil
}
