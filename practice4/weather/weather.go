package weather

type DailyWeather struct {
	Weather []WeatherForecast `json:"list"`
	City    struct {
		Name    string `json:"name"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"city"`
}

type WeatherForecast struct {
	Dt   int64 `json:"dt"`
	Main struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Sky []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  int     `json:"gust"`
	} `json:"wind"`
	DtTxt string `json:"dt_txt"`
}
