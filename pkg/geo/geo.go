package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponce struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		i, _ := checkCity(city)
		if i == true {
			return &GeoData{City: city}, nil
		} else {
			panic("Не существующий город")
		}

	}
	resp, err := http.Get("http://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil

}

func checkCity(city string) (bool, error) {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("http://countriesnov.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false, errors.New("NOT200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var populationResponse CityPopulationResponce
	err = json.Unmarshal(body, &populationResponse)
	if err != nil {
		return false, err
	}
	return populationResponse.Error, nil

}
