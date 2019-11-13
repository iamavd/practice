package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Meteorologist struct {
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

func (w Weather) GetWind() (speed float64, direction string) {
	speed = w.Wind.Speed
	direction = windDirection(w.Wind.Deg)
	return
}

func windDirection(deg int) string {
	switch {
	case deg < 45:
		return "СВ"
	case deg > 45, deg < 90:
		return "В"
	case deg > 90, deg < 135:
		return "ЮВ"
	case deg > 135, deg < 180:
		return "Ю"
	case deg > 180, deg < 225:
		return "ЮЗ"
	case deg > 225, deg < 270:
		return "З"
	case deg > 270, deg < 315:
		return "СЗ"
	case deg > 315:
		return "С"
	}
	return "--"
}

/*
func (m Meteorologist)PrintForecast() {
	fmt.Printf("Сегодня в городе %v %v, температура воздуха %v°С, ветер %v%м/с с порывами до %vм/с. Влажность воздуха %v%%. Восход солнца %v, заход солнца %v.", m.WeatherForecast.)
}*/
