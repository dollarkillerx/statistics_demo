package api

import (
	"github.com/dollarkillerx/common/pkg/open_telemetry"
	"github.com/gin-gonic/gin"
)

func (a *ApiServer) Router() {
	a.app.GET("/health", a.HealthCheck)
}

func (a *ApiServer) HealthCheck(ctx *gin.Context) {
	_, span := open_telemetry.Tracer.Start(ctx, ctx.Request.URL.Path)
	defer span.End()

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
