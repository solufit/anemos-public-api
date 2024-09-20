package version

import "github.com/gin-gonic/gin"

type VersionHandler struct{}

type Response struct {
	Status string `json:"status"`
}

// Version godoc
// @Summary バージョン管理
// @Tags version
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/version [get]
func (h *VersionHandler) V1Version(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Rerease": "2024-09-19",
	})
}

func NewVersionHandler() *VersionHandler {
	return &VersionHandler{}
}
