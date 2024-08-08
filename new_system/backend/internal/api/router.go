package api

import (
	"github.com/dollarkillerx/common/pkg/open_telemetry"
	"github.com/gin-gonic/gin"
)

func (a *ApiServer) Router() {
	a.app.GET("/health", a.HealthCheck)

	ea := a.app.Group("/ea")
	{
		// broadcast 广播
		ea.POST("/broadcast", a.broadcast)
		// subscription 订阅
		ea.POST("/subscription", a.subscription)
		// errors 错误
		ea.POST("/errors", a.errors)
	}
	api := a.app.Group("/api")
	{
		// accounts 账户
		api.GET("/accounts", a.accounts)
		// account 账户的信息
		api.GET("/account/:account", a.account)
		api.GET("/account/charts/:account", a.chartsAccount)
	}
}

func (a *ApiServer) HealthCheck(ctx *gin.Context) {
	_, span := open_telemetry.Tracer.Start(ctx, ctx.Request.URL.Path)
	defer span.End()

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
