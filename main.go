package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const openWeatherMapAPIKey = "983dc274aec42fb5bb52778a7d2805d3"
const city = "Toronto"

type WeatherData struct {
	City        string  `json:"city"`
	Temperature string  `json:"temperature"`
	Description string  `json:"description"`
	Humidity    float64 `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	WeatherIcon string  `json:"weather_icon"`
}

func kelvinToCelsius(kelvin float64) string {
	celsius := kelvin - 273.15
	return fmt.Sprintf("%.2f", celsius)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	port := 7575
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
