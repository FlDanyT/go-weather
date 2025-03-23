package weather

import (

	"demo/weather/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

)

var ErrWrongFormat = errors.New("WRONG_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {

	if format < 0 || format > 4 {

		return "", ErrWrongFormat

	}

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City) //  Разбирает строку URL и преобразует ее в структуру url.URL

	if err != nil {

		fmt.Println(err.Error())
		return "", errors.New("ERROR__URL")

	}

	params := url.Values{}

	params.Add("format", fmt.Sprint(format)) // Добавляем format в ссылку запроса
	baseUrl.RawQuery = params.Encode() // Выполняет URL-кодирование специальных символов

	resp, err := http.Get(baseUrl.String())

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR__HTTP")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		fmt.Println(err.Error())
		return "", errors.New("ERROR__READBODY")
	}

	return string(body), nil // Отдаем ответ не в байтах а не в строке

}