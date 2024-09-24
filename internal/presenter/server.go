package presenter

import (
	"context"
	version "solufit/go/internal/service/version"
	weather "solufit/go/internal/service/weather" // Add this import

	"github.com/gin-gonic/gin"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// バージョン管理
	{
		versionHandler := version.NewVersionHandler()
		v1.GET("/version", versionHandler.V1Version)
	}
	{
		weatherHandler := weather.NewWeatherHandler()
		v1.POST("/weather/info", weatherHandler.PostWeatherInfoNear)
		v1.POST("/weather/info/past", weatherHandler.PostWeatherInfoPast)
	}

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
