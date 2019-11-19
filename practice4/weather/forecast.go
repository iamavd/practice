package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const ansoluteZero = -273.15
const apiKey = "2c19a8c670afc70f2ae7a81f229fce3d"

type Meteorologist struct {
	city string
}

func (m Meteorologist) WeatherForecast(city string, cnt string) DailyWeather {
	dw := DailyWeather{}
	link := createLink(city, cnt)
	resp := weatherResponse(link)
	err := json.Unmarshal(resp, &dw)
	if err != nil {
		fmt.Println("Error")
	}
	return dw
}

func weatherResponse(link string) []byte {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body
}

func createLink(city string, cnt string) string {
	return "http://api.openweathermap.org/data/2.5/forecast?q=" + city + "&lang=ru&cnt=" + cnt + "&appid=" + apiKey
}

func (w WeatherForecast) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	temp = w.Main.Temp + ansoluteZero
	tempMin = w.Main.TempMin + ansoluteZero
	tempMax = w.Main.TempMax + ansoluteZero
	return
}

func (w WeatherForecast) GetCloudiness() string {
	return w.Sky[0].Description
}

func (w WeatherForecast) GetHumidity() int {
	return w.Main.Humidity
}

func (w WeatherForecast) GetWind() (speed float64, direction string, gust int) {
	speed = w.Wind.Speed
	direction = windDirection(w.Wind.Deg)
	gust = w.Wind.Gust
	return
}

func windDirection(deg int) string {
	switch {
	case deg > 45, deg < 135:
		return "восточный"
	case deg > 135, deg < 225:
		return "южный"
	case deg > 225, deg < 315:
		return "западный"
	case deg > 315, deg < 45:
		return "северный"
	}
	return ""
}

func PrintForecast(weather DailyWeather) {
	for _, w := range weather.Weather {

		temp, _, _ := w.GetTemperature()
		speed, dir, gust := w.GetWind()
		gustOut := ""
		if gust != 0 {
			gustOut = fmt.Sprintf(",c порывами до %v м/c", gust)
		}
		fmt.Printf("%v в городе %v %v, температура воздуха %.1f°С, ветер %v %v м/с%v.Влажность воздуха %v%%.\n", printDate(w.Dt), weather.City.Name, w.GetCloudiness(), temp, dir, speed, gustOut, w.GetHumidity())
	}
	fmt.Printf("Восход солнца %v, заход солнца %v", encodeTime(weather.City.Sunrise), encodeTime(weather.City.Sunset))
}

func encodeTime(utc int64) string {
	return time.Unix(utc, 0).Format("15:04")
}

func printDate(utc int64) string {
	dt := time.Unix(utc, 0)
	switch dt.Day() {
	case time.Now().Day():
		return "Сегодня " + dt.Format("15:04")
	case time.Now().AddDate(0, 0, 1).Day():
		return "Завтра " + dt.Format("15:04")
	default:
		return dt.Format("02.01.2006 15:04")
	}
	return ""
}
