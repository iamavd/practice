package main

import (
	"practice4/weather"
)

func main() {
	m := weather.Meteorologist{}
	weather.PrintForecast(m.WeatherForecast("Moscow", "20"))
}
