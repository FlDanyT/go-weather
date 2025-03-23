package weather_test

import (

	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"

)

func TestGetWeather(t *testing.T) {

	// Arrange
	expected := "London"
	geoData := geo.GeoData{

		City: expected,
	}

	format := 3

	// Act
	result, err := weather.GetWeather(geoData, format)

	// Assert
	if err != nil {

		t.Errorf("Пришла ошибка %v", err)

	}

	if !strings.Contains(result, expected) {

		t.Errorf("Ожидалось %v, получение %v", expected, result)

	}

}

var testCases = []struct { // Группа тестов
	name   string
	format int
}{

	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestWeatherWrongFormat(t *testing.T) { // Негативный тест, тестируем неверный формат

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			expected := "London"
			geoData := geo.GeoData{

				City: expected,
			}

			// Act
			_, err := weather.GetWeather(geoData, tc.format)

			// Assert
			if err != weather.ErrWrongFormat {

				t.Errorf("Ожидалось %v, получение %v", weather.ErrWrongFormat, err)

			}

		})

	}

}