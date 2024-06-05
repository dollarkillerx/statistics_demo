package api

import "github.com/gin-gonic/gin"

func (a *ApiServer) RegisterRoutes() {
	v1 := a.gin.Group("/api/v1")
	{
		v1.POST("/init", a.apiInit)
		v1.POST("/symbol_info", a.symbolInfo)
		v1.POST("/symbol_info_tick", a.symbolInfoTick)
		v1.POST("/order_send", a.orderSend)
		v1.POST("/positions_total", a.positionsTotal)
		v1.POST("/positions_get", a.positionsGet)
		v1.POST("/account_info", a.accountInfo)
		v1.POST("/account_info", a.accountInfo)
	}
}

func (a *ApiServer) apiInit(c *gin.Context) {

}

func (a *ApiServer) symbolInfo(c *gin.Context) {

}

func (a *ApiServer) symbolInfoTick(c *gin.Context) {

}

func (a *ApiServer) orderSend(c *gin.Context) {

}

func (a *ApiServer) positionsTotal(c *gin.Context) {

}

func (a *ApiServer) positionsGet(c *gin.Context) {

}

func (a *ApiServer) accountInfo(c *gin.Context) {

}
