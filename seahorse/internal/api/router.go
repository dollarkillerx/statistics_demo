package api

import (
	"github.com/gin-gonic/gin"
	"math"
	"seahorse/internal/models"
)

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
	var req models.ReqInit
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	err := a.storage.Bb.Model(&models.Account{}).Create(&models.Account{
		Account:         req.Account,
		Balance:         req.Balance,
		Lever:           req.Lever,
		LargestPosition: 0,
		LargestLoss:     0,
		LargestProfit:   0,
	}).Error
	if err != nil {
		panic(err)
	}

	c.JSON(200, models.RespInit{
		Account: req.Account,
		Error:   "",
	})
}

func (a *ApiServer) symbolInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"point": 0.00001,
	})
}

func (a *ApiServer) symbolInfoTick(c *gin.Context) {
	var req models.ReqSymbolInfoTick
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	tick := a.storage.GetTick()
	c.JSON(200, models.RespSymbolInfoTick{
		Ask:       tick.Ask,
		Bid:       tick.Bid,
		Timestamp: tick.Timestamp,
	})

}

func (a *ApiServer) orderSend(c *gin.Context) {
	var req models.ReqOrderSend
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	// 未来考虑加入滑点

	tick2 := a.storage.GetTick2()

	// 进场
	if req.Position == 0 {
		err := a.storage.Bb.Model(&models.Order{}).Create(&models.Order{
			Symbol:     req.Symbol,
			Type:       req.Type,
			Volume:     req.Volume,
			CreateTime: tick2.Timestamp,
			Price:      req.Price,
			CloseTime:  0,
			Profit:     0,
		}).Error
		if err != nil {
			panic(err)
		}
	} else {
		// 出场
		var order models.Order
		err := a.storage.Bb.Model(&models.Order{}).Where("symbol = ? and id = ?", req.Symbol, req.Position).First(&order).Error
		if err != nil {
			panic(err)
		}

		// 计算盈利
		var profit float64
		if req.Type == 0 {
			profit = math.Round(((order.Price-req.Price)*req.Volume*100000)*100) / 100
		} else {
			profit = math.Round(((req.Price-order.Price)*req.Volume*100000)*100) / 100
		}

		err = a.storage.Bb.Model(&models.Order{}).Where("symbol = ? and id = ?", req.Symbol, req.Position).Updates(&models.Order{
			ClosePrice: req.Price,
			CloseTime:  tick2.Timestamp,
			Profit:     profit,
		}).Error
		if err != nil {
			panic(err)
		}
	}

	c.JSON(200, gin.H{
		"error": "",
	})
}

func (a *ApiServer) positionsTotal(c *gin.Context) {
	var count int64
	err := a.storage.Bb.Model(&models.Order{}).Where("close_time = 0").Count(&count).Error
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"total": count,
		"error": "",
	})
}

func (a *ApiServer) positionsGet(c *gin.Context) {

}

func (a *ApiServer) accountInfo(c *gin.Context) {

}
