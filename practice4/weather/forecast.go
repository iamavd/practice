package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Meteorologist struct {
	city string
}

func (m Meteorologist) WeatherForecast(city string) Weather {
	w := Weather{}
	link := createLink(city)
	resp := weatherResponse(link)
	err := json.Unmarshal(resp, &w)
	if err != nil {
		fmt.Println("Error")
	}
	return w
}

func weatherResponse(link string) []byte {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body
}

func createLink(city string) string {
	return "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d"
//	return "http://api.openweathermap.org/data/2.5/forecast?q=Moscow,us&appid=2c19a8c670afc70f2ae7a81f229fce3d" //link for 5 day forecast
}

func (w Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	temp = w.Main.Temp
	tempMin = w.Main.TempMin
	tempMax = w.Main.TempMax
	return
}

func (w Weather) GetCloudiness() string {
	return w.Sky[0].Description
}

func (w Weather) GetHumidity() int {
	return w.Main.Humidity
}

func (w Weather) GetWind() (speed float64, direction string, gust int) {
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

func PrintForecast(w Weather) {
	temp, _, _ := w.GetTemperature()
	speed, dir, gust := w.GetWind()
	gustOut := ""
	if gust != 0 {
		gustOut = fmt.Sprintf(",c порывами до %v м/c",gust)
	}
	fmt.Printf("Сегодня в городе %v %v, температура воздуха %.1f°С, ветер %v %v м/с%v.Влажность воздуха %v%%. Восход солнца %v, заход солнца %v","CITYNAME",w.GetCloudiness(),temp,dir,speed,gustOut,w.GetHumidity(),encodeTime(w.Sys.Sunrise), encodeTime(w.Sys.Sunset)) //, ветер %v%м/с с порывами до %vм/с. Влажность воздуха %v%%. Восход солнца %v, заход солнца %v.
}

func encodeTime(utc int64) string{
	return time.Unix(utc,0).Format("15:04")
}
