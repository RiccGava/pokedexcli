package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"

	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("FOUND CACHE")
		locationAreaResp := LocationAreaResp{}

		err := json.Unmarshal(data, &locationAreaResp)

		if err != nil {
			return LocationAreaResp{}, err
		}

		return locationAreaResp, nil
	}
	fmt.Println("CACHE MISS")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}

	err = json.Unmarshal(data, &locationAreaResp)

	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, data)
	return locationAreaResp, nil

}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName

	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("FOUND CACHE")
		locationArea := LocationArea{}

		err := json.Unmarshal(data, &locationAreaName)

		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}
	fmt.Println("CACHE MISS")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}

	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)
	return locationArea, nil

}
