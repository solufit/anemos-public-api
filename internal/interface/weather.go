package weather_interface

type WeatherInfoFilter struct {
	Number    int    `json:"number"`
	PostCode  string `json:"post_code"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Filter    string `json:"filter" enum:"weather-forecast,weather-warning,earthquake"`
}
