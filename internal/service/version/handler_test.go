package version

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestV1Version(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	NewVersionHandler().V1Version(ctx)
	assert.Equal(t, 200, ctx.Writer.Status())

}
