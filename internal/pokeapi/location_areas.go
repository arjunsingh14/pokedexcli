package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

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
	var locationAreas locationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreas)
	if err != nil {
		return locationArea{}, err
	}
	body, _ := io.ReadAll(res.Body)
	c.cache.Add(url, body)
	return locationAreas, nil
}