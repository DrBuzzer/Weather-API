package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/weather", weatherHandler)
	http.ListenAndServe(":5555", nil)
}

type Weather struct {
	Coord       Coord         `json:"coord"`
	WeatherInfo []WeatherInfo `json:"weather"`
	Base        string        `json:"base"`
	Main        Main          `json:"main"`
	Visibility  int           `json:"visibility"`
	Wind        Wind          `json:"wind"`
	Clouds      Clouds        `json:"clouds"`
	Dt          int64         `json:"dt"`
	Sys         Sys           `json:"sys"`
	Timezone    int           `json:"timezone"`
	ID          int           `json:"id"`
	Name        string        `json:"name"`
}
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type WeatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Ваш код для выполнения запроса к внешнему API погоды здесь
	apiKey := "8ad23f027139c12aef619242f1d2283f"
	city := r.URL.Query().Get("city")
	safeCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", safeCity, apiKey)

	// Выполнение запроса к внешнему API погоды
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Обработка ответа от внешнего API погоды
	var weatherData Weather
	err = json.NewDecoder(response.Body).Decode(&weatherData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Сериализация структуры Weather в формат JSON
	jsonResponse, err := json.Marshal(weatherData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Установка заголовков для ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Отправка ответа в формате JSON
	w.Write(jsonResponse)
}
