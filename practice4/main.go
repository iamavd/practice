package main

import (
	"fmt"
	"practice4/weather"
)

func main() {
	//weather.Parser()
	m := weather.Meteorologist{}
	fmt.Println(m.WeatherForecast("Moscow").GetCloudiness())

	weather.PrintForecast(m.WeatherForecast("Mahilyow"))
}
