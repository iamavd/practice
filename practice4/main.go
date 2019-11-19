package main

import (
	"practice4/weather"
)

func main() {
	m := weather.Meteorologist{}

	m.WeatherForecast("Moscow").PrintForecast()
	m.DailyForecast("Mahilyow", 20).PrintForecast()

}
