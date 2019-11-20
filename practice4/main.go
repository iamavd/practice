package main

import (
	"practice4/weather"
	"fmt"
	"os"
)

func main() {
	m := weather.Meteorologist{}

	today, err := m.WeatherForecast("Moscow")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	daily, err := m.DailyForecast("Mahilyow", 20)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	today.PrintForecast()
	daily.PrintForecast()
}
