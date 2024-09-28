package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiKey = "20704c12dbd55c4cde366d56a68eac21"

/* Sample return
{
    "coord": {
        "lon": -0.1257,
        "lat": 51.5085
    },
    "weather": [
        {
            "id": 804,
            "main": "Clouds",
            "description": "overcast clouds",
            "icon": "04n"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 283.7,
        "feels_like": 282.77,
        "temp_min": 281.55,
        "temp_max": 284.87,
        "pressure": 1026,
        "humidity": 75,
        "sea_level": 1026,
        "grnd_level": 1022
    },
    "visibility": 10000,
    "wind": {
        "speed": 1.38,
        "deg": 255,
        "gust": 2.46
    },
    "clouds": {
        "all": 100
    },
    "dt": 1727550636,
    "sys": {
        "type": 2,
        "id": 2075535,
        "country": "GB",
        "sunrise": 1727503000,
        "sunset": 1727545520
    },
    "timezone": 3600,
    "id": 2643743,
    "name": "London",
    "cod": 200
}
*/

type Main struct {
	Temp float64 `json:"temp"`
}

type Data struct {
	MainField Main `json:"main"`
}

func fetchWeather(city string) (Data, error) {

	data := Data{}

	// https://openweathermap.org/current , https://home.openweathermap.org/api_keys
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data, err
	}

	defer resp.Body.Close()

	// my way
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Http response failed %s\n", err)
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error decoding weather data for %s\n", err)
		return data, err
	}

	// another way
	// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
	// 	fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
	// 	return data
	// }

	// ch <- data

	return data, nil
}

func fetchWeather2(city string, ch chan<-Data, chErr chan<-error) {

	data := Data{}

	// https://openweathermap.org/current , https://home.openweathermap.org/api_keys
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		chErr<- err
	}

	defer resp.Body.Close()

	// my way
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		chErr<- err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		chErr<- err
	}

	// another way
	// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
	// 	fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
	// 	return data
	// }

	ch <- data
}

func main() {
	startNow := time.Now()

	ch := make(chan Data, 5)
	chErr := make(chan error)

	cities := []string{"toronto","london", "paris", "tokyo", "beijing"}

	// for _, city := range cities {
	// 	data, _ := fetchWeather(city)
	// 	fmt.Printf("This is the temp %v from %s\n", data.MainField.Temp, city)
	// }

	data := Data{}
	err := errors.New("")
	for _, city := range cities {
		go fetchWeather2(city, ch, chErr)
	}

	for i := 0; i < len(cities); i++ {
		select {
		case err = <-chErr:
			fmt.Printf("some err: %s", err.Error())
		case data = <-ch:
			fmt.Printf("This is the temp %v from %s\n", data.MainField.Temp, cities[i])
		}
	}

	close(ch)
	close(chErr)

	fmt.Println("This operation took: ", time.Since(startNow))
}