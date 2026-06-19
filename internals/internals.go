package internals

import (
	"encoding/json"
	"net/http"
)

func GetLocationAreas(url string) (locationArea, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
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

	return locationAreas, nil
}