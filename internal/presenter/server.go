package presenter

import (
	"context"
	"solufit/go/internal/service/version"

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
		v1.GET("/v1/version", versionHandler.V1Version)
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
