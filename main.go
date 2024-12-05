package main

import (
	"flag"
	"fmt"
	"local/pkg/geo"
	"local/pkg/geo/weather"
)

func main() {
	city := flag.String("city", "", "Users city")
	format := flag.Int("format", 1, "Weather output format")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(weatherData)

}
