package weather

type Weather struct {
	Sky []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  int `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Sys struct {
		Sunrise int64    `json:"sunrise"`
		Sunset  int64    `json:"sunset"`
	} `json:"sys"`
	SunriseDaily    int    `json:"sunrise"`
	SunsetDaily     int    `json:"sunset"`
}

type DailyWeather struct {
}
