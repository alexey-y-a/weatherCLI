package main

import (
	"flag"
	"fmt"
	"weathercli/geo"
	"weathercli/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Printf("%s\n %s\n", geoData.City, weatherData)
}
