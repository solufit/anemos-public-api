package version

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostWeatherInfoPast(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "valid request",
			requestBody:  `{"number":1,"post_code":"123-4567","start_date":"2024-09-18","end_date":"2024-09-19","filter":"weather-forecast"}`,
			expectedCode: http.StatusOK,
			expectedBody: `[{"id":"1","object_type":"weather","areacode":"123-4567","title":"気象情報","status":"success","detail":{"weather_today":"晴れ","weather_tomorrow":"晴れ","max_temp":30,"min_temp":20,"rain_percent_now":0,"rain_percent_6h":0,"rain_percent_12h":0,"rain_percent_18h":0,"rain_percent_24h":0},"reported_at":"2024-09-19","info_domain":"weather","info_objectId":"1"}]`,
		},
		{
			name:         "invalid request",
			requestBody:  `{"number":1,"start_date":"2024-09-18","end_date":"2024-09-19","filter":"invalid-filter"}`,
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: `[{"status":"Validation error -> Key: 'WeatherInfoFilter.Filter' Error:Field validation for 'Filter' failed on the 'enum' tag"}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			handler := NewWeatherHandler()
			router.POST("/v1/weather/info/past", handler.PostWeatherInfoPast)

			req, _ := http.NewRequest(http.MethodPost, "/v1/weather/info/past", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestPostWeatherInfoNear(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "valid request",
			requestBody:  `{"number":1,"post_code":"123-4567","start_date":"2024-09-18","end_date":"2024-09-19","filter":"weather-forecast"}`,
			expectedCode: http.StatusOK,
			expectedBody: `[{"id":"1","object_type":"weather","areacode":"123-4567","title":"気象情報","status":"success","detail":{"weather_today":"晴れ","weather_tomorrow":"晴れ","max_temp":30,"min_temp":20,"rain_percent_now":0,"rain_percent_6h":0,"rain_percent_12h":0,"rain_percent_18h":0,"rain_percent_24h":0},"reported_at":"2024-09-19","info_domain":"weather","info_objectId":"1"}]`,
		},
		{
			name:         "invalid request",
			requestBody:  `{"number":1,"post_code":"123-4567","start_date":"2024-09-18","end_date":"2024-09-19","filter":"invalid-filter"}`,
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: `{"status":"Validation errorKey: 'WeatherInfoFilter.Filter' Error:Field validation for 'Filter' failed on the 'enum' tag"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			handler := NewWeatherHandler()
			router.POST("/v1/weather/info", handler.PostWeatherInfoPast)

			req, _ := http.NewRequest(http.MethodPost, "/v1/weather/info/past", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
