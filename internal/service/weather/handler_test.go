package version

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestV1GetWeatherInfoNear(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.BindJSON(&WeatherInfoFilter{
		Number:    1,
		PostCode:  "123-4567",
		StartDate: "2024-09-19",
		EndDate:   "2024-09-19",
		Filter:    "weather-forecast",
	})
	NewWeatherHandler().GetWeatherInfoNear(ctx)
	assert.Equal(t, 200, ctx.Writer.Status())
}

func TestV1GetWeatherInfoPast(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.BindJSON(&WeatherInfoFilter{
		Number:    1,
		PostCode:  "123-4567",
		StartDate: "2024-09-19",
		EndDate:   "2024-09-19",
		Filter:    "weather-forecast",
	})
	NewWeatherHandler().GetWeatherInfoPast(ctx)
	fmt.Println(ctx.Writer.Status())
	assert.Equal(t, 200, ctx.Writer.Status())
}
