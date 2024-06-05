package api

import (
	"math"

	"github.com/gin-gonic/gin"
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

	a.gin.GET("web", a.web)
}

func (a *ApiServer) web(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
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
	var req models.ReqOrderPositionsGet
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	var orders []models.Order
	err := a.storage.Bb.Model(&models.Order{}).Where("close_time = 0").Order("create_time").Find(&orders).Error
	if err != nil {
		panic(err)
	}

	tick2 := a.storage.GetTick2()

	var items []models.RespOrderPosition
	for _, order := range orders {
		price := tick2.Ask
		var profile float64
		if order.Type == 1 { // 如果当前订单为sell
			price = tick2.Ask // 这做空平仓
			profile = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
		} else {
			price = tick2.Bid // 做多
			profile = math.Round(((price-order.Price)*order.Volume*100000)*100) / 100
		}

		items = append(items, models.RespOrderPosition{
			Ticket:       order.ID,
			Time:         order.CreateTime,
			Type:         order.Type,
			Volume:       order.Volume,
			PriceOpen:    order.Price,
			PriceCurrent: price,
			Profit:       profile,
		})
	}

	c.JSON(200, models.RespOrderPositionsGet{
		Items: items,
	})
}

func (a *ApiServer) accountInfo(c *gin.Context) {
	var orders []models.Order
	err := a.storage.Bb.Model(&models.Order{}).Where("close_time = 0").Order("create_time").Find(&orders).Error
	if err != nil {
		panic(err)
	}

	tick2 := a.storage.GetTick2()

	var items []models.RespOrderPosition
	for _, order := range orders {
		price := tick2.Ask
		var profile float64
		if order.Type == 1 { // 如果当前订单为sell
			price = tick2.Ask // 这做空平仓
			profile = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
		} else {
			price = tick2.Bid // 做多
			profile = math.Round(((price-order.Price)*order.Volume*100000)*100) / 100
		}

		items = append(items, models.RespOrderPosition{
			Ticket:       order.ID,
			Time:         order.CreateTime,
			Type:         order.Type,
			Volume:       order.Volume,
			PriceOpen:    order.Price,
			PriceCurrent: price,
			Profit:       profile,
		})
	}

	var profit float64
	for _, item := range items {
		profit += item.Profit
	}

	c.JSON(200, models.RespAccountInfo{
		Profit: profit,
	})
}
