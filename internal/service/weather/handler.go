package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct{}

type Response struct {
	Status string `json:"status"`
}

type WeatherInfoFilter struct {
	Number    int    `json:"number"`
	PostCode  string `json:"post_code"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Filter    string `json:"filter" enum:"weather-forecast,weather-warning,earthquake"`
}

type DetailEartquake struct {
	EventID          string `json:"event_id"`
	EntryID          string `json:"entry_id"`
	EditorialOffice  string `json:"editorial_id"`
	PublishingOffice string `json:"publishing_office"`
	Category         string `json:"category"`
	Datetime         string `json:"datetime"`
	Headline         string `json:"headline"`
	Hypocenter       string `json:"hypocenter"`
	RegionCode       string `json:"region_code"`
	MaxInt           string `json:"max_int"`
	Magnitude        string `json:"magnitude"`
}

type DetailWeatherWarning struct {
	EntryID          string `json:"entry_id"`
	EditorialOffice  string `json:"editorial_id"`
	PublishingOffice string `json:"publishing_office"`
	Category         string `json:"category"`
	Datetime         string `json:"datetime"`
	Headline         string `json:"headline"`
	Pref             string `json:"pref"`
}

type DetailWeatherForecast struct {
	WeatherToday    string `json:"weather_today"`
	WeatherTomorrow string `json:"weather_tomorrow"`
	MaxTemp         int    `json:"max_temp"`
	MinTemp         int    `json:"min_temp"`
	RainPercentNow  int    `json:"rain_percent_now"`
	RainPercent6h   int    `json:"rain_percent_6h"`
	RainPercent12h  int    `json:"rain_percent_12h"`
	RainPercent18h  int    `json:"rain_percent_18h"`
	RainPercent24h  int    `json:"rain_percent_24h"`
}

type WeatherInfo struct {
	ID           string      `json:"id"`
	ObjectType   string      `json:"object_type"`
	Areacode     string      `json:"areacode"`
	Title        string      `json:"title"`
	Status       string      `json:"status"`
	Detail       interface{} `json:"detail"`
	ReportedAt   string      `json:"reported_at"`
	InfoDomain   string      `json:"info_domain"`
	InfoObjectId string      `json:"info_objectId"`
}

// Version godoc
// @Summary 気象情報
// @Tags weather
// @Accept json
// @Produce json
// @Success 200 {object} WeatherInfo
// @Router /v1/version [get]
func (h *WeatherHandler) GetWeatherInfoNear(ctx *gin.Context) {
	var filter WeatherInfoFilter
	if err := ctx.BindJSON(&filter); err != nil {
		return
	}

	var weatherinfo = []WeatherInfo{
		{
			ID:         "1",
			ObjectType: "weather",
			Areacode:   "123-4567",
			Title:      "気象情報",
			Status:     "success",
			Detail: DetailWeatherForecast{
				WeatherToday:    "晴れ",
				WeatherTomorrow: "晴れ",

				MaxTemp:        30,
				MinTemp:        20,
				RainPercentNow: 0,
				RainPercent6h:  0,
				RainPercent12h: 0,
				RainPercent18h: 0,
				RainPercent24h: 0,
			},
			ReportedAt:   "2024-09-19",
			InfoDomain:   "weather",
			InfoObjectId: "1",
		},
	}

	ctx.IndentedJSON(http.StatusOK, weatherinfo)
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}
