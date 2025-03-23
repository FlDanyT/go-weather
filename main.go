package main

import (

	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"

)

func main() {

	fmt.Println("Передайте flag с городом")
	
	city := flag.String("city", "", "Город пользователя") // Передаем информацию через строку при запуске
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city) // Получаем место положение

	if err != nil {

		fmt.Println(err.Error())

	}

	weatherData, _ := weather.GetWeather(*geoData, *format) // Получаем место положение города
	fmt.Println(weatherData)

}
