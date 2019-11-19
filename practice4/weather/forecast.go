package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const ansoluteZero = -273.15
const apiKey = "2c19a8c670afc70f2ae7a81f229fce3d"

type Meteorologist struct {
}

func (m Meteorologist) DailyForecast(city string, cnt int) DailyWeather {
	dw := DailyWeather{}
	link := createLink(city, cnt)
	resp := weatherResponse(link)
	err := json.Unmarshal(resp, &dw)
	if err != nil {
		fmt.Println("Error")
	}
	return dw
}

func(m Meteorologist) WeatherForecast(city string) DailyWeather {
	return m.DailyForecast(city,1)
}
func weatherResponse(link string) []byte {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body
}

func createLink(city string, cnt int) string {
	return "http://api.openweathermap.org/data/2.5/forecast?q=" + city + "&lang=ru&cnt=" + strconv.Itoa(cnt) + "&appid=" + apiKey
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

func (weather DailyWeather) PrintForecast() {
	for _, w := range weather.Weather {
		temp, _, _ := w.GetTemperature()
		speed, dir, gust := w.GetWind()
		gustOut := ""
		if gust != 0 {
			gustOut = fmt.Sprintf(",c порывами до %v м/c", gust)
		}
		fmt.Printf("%v в городе %v %v, температура воздуха %.1f°С, ветер %v %v м/с%v.Влажность воздуха %v%%.\n", printDate(w.Dt), weather.City.Name, w.GetCloudiness(), temp, dir, speed, gustOut, w.GetHumidity())
	}
	fmt.Printf("Восход солнца %v, заход солнца %v\n------\n", encodeTime(weather.City.Sunrise), encodeTime(weather.City.Sunset))
}