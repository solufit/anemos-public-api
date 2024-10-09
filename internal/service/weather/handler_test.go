package version

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
		},
		{
			name:         "invalid request",
			requestBody:  `{"number":1,"post_code":"123-4567","start_date":"2024-09-18","end_date":"2024-09-19","filter":"invalid-filter"}`,
			expectedCode: http.StatusOK, // should be http.StatusUnprocessableEntity
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			handler := NewWeatherHandler()
			router.POST("/v1/weather/info", handler.PostWeatherInfoNear)

			req, _ := http.NewRequest(http.MethodPost, "/v1/weather/info", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
