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

func GetMyLocation(city string) (*GeoData, error) { // Наша геолокация

	if city != "" {

		isCity := checkCity(city)
		if !isCity {

			panic("Такого города нет")

		}

		return &GeoData{
			City: city,
		}, nil

	}

	resp, err := http.Get("http://ipapi.co/json/") // Делаем запрос
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {

		return nil, errors.New("NOT200")

	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Читает данные в байтах
	if err != nil {

		return nil, err

	}

	var geo GeoData
	json.Unmarshal(body, &geo) // Извлекает данные из json в батах и записываем в формате go в geo
	return &geo, nil

}

func checkCity(city string) bool { // Проверка есть ли такой город

	postBody, _ := json.Marshal(map[string]string{ // Body
		"city": city,
	})

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))

	if err != nil {

		return false

	}

	defer resp.Body.Close() // Закрываем чтение Body

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return false

	}

	var populationResponce CityPopulationResponce

	json.Unmarshal(body, &populationResponce) // Извлекает данные из json в батах и записываем в формате go в geo
	return !populationResponce.Error

}