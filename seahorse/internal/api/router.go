package api

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"seahorse/internal/models"
)

func (a *ApiServer) RegisterRoutes() {
	v1 := a.gin.Group("/api/v1")
	{
		v1.POST("/init", a.apiInit)
		v1.POST("/symbol_info", a.symbolInfo)
		v1.POST("/symbol_info_tick", a.symbolInfoTick)
		v1.POST("/symbol_info_tick2", a.symbolInfoTick2)
		v1.POST("/order_send", a.orderSend)
		v1.POST("/positions_total", a.positionsTotal)
		v1.POST("/positions_get", a.positionsGet)
		v1.POST("/close_all", a.closeAll)
		v1.POST("/account_info", a.accountInfo)
	}

	a.gin.GET("web", a.web)
}

func (a *ApiServer) apiInit(c *gin.Context) {
	var req models.ReqInit
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	err := a.storage.Bb.Model(&models.Account{}).Create(&models.Account{
		Account:         req.Account,
		InitialAmount:   req.Balance,
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

	// 统计tp sl
	var orders []models.Order
	a.storage.Bb.Model(&models.Order{}).
		Where("close_time = 0").Find(&orders)
	for _, order := range orders {
		if order.Tp == 0 && order.Sl == 0 {
			continue
		}
		// buy
		if order.Type == 0 {
			if order.Tp != 0 && tick.Bid >= order.Tp {
				a.closeOrder(order, tick)
			}

			if order.Sl != 0 && tick.Bid <= order.Sl {
				a.closeOrder(order, tick)
			}
		}
		// sell
		if order.Type == 1 {
			if order.Tp != 0 && tick.Ask <= order.Tp {
				a.closeOrder(order, tick)
			}

			if order.Sl != 0 && tick.Ask >= order.Sl {
				a.closeOrder(order, tick)
			}
		}
	}

	c.JSON(200, models.RespSymbolInfoTick{
		Ask:       tick.Ask,
		Bid:       tick.Bid,
		Timestamp: tick.Timestamp,
		Time:      tick.Timestamp,
	})
}

func (a *ApiServer) closeOrder(order models.Order, tick2 models.Tick) {
	// 计算盈利
	var profit float64
	var price float64
	if order.Type == 1 { // 如果当前订单为sell
		price = tick2.Ask // 这做空平仓
		profit = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
	} else {
		price = tick2.Bid // 做多
		profit = math.Round(((price-order.Price)*order.Volume*100000)*100) / 100
	}

	err := a.storage.Bb.Model(&models.Order{}).Where("id = ?", order.ID).Updates(&models.Order{
		ClosePrice:   price,
		CloseTime:    tick2.Timestamp,
		CloseTimeStr: time.Unix(tick2.Timestamp, 0).Format("2006-01-02 15:04:05"),
		Profit:       profit,
		Auto:         true,
	}).Error
	if err != nil {
		panic(err)
	}

	a.storage.Bb.Model(&models.Account{}).Where("account = ?", order.Account).UpdateColumn("balance", gorm.Expr("balance + ?", profit))
}

func (a *ApiServer) symbolInfoTick2(c *gin.Context) {
	var req models.ReqSymbolInfoTick
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	tick := a.storage.GetTick2()
	c.JSON(200, models.RespSymbolInfoTick{
		Ask:       tick.Ask,
		Bid:       tick.Bid,
		Timestamp: tick.Timestamp,
		Time:      tick.Timestamp,
	})
}

func (a *ApiServer) orderSend(c *gin.Context) {
	var req models.ReqOrderSend
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	// 未来考虑加入滑点
	account := a.storage.GetAccount(req.Account)
	if account.Profit < 0 {
		if account.Profit+account.Balance-account.Margin <= account.Balance {
			c.JSON(500, gin.H{
				"error": "Liquidation爆仓",
			})
			return
		}
	}

	tick2 := a.storage.GetTick2()

	// 进场
	if req.Position == 0 {
		err := a.storage.Bb.Model(&models.Order{}).Create(&models.Order{
			Account:       req.Account,
			Symbol:        req.Symbol,
			Type:          req.Type,
			Volume:        req.Volume,
			CreateTime:    tick2.Timestamp,
			CreateTimeStr: time.Unix(tick2.Timestamp, 0).Format("2006-01-02 15:04:05"),
			Price:         req.Price,
			CloseTime:     0,
			Profit:        0,
			Tp:            req.Tp,
			Sl:            req.Sl,
			Comment:       req.Comment,
			Margin:        req.Price * req.Volume * 100000 / float64(account.Lever),
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
		var price float64
		if order.Type == 1 { // 如果当前订单为sell
			price = tick2.Ask // 这做空平仓
			profit = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
		} else {
			price = tick2.Bid // 做多
			profit = math.Round(((price-order.Price)*order.Volume*100000)*100) / 100
		}

		if req.Price != price {
			log.Println("req: ", req.Price)
			log.Println("pr: ", price)
			log.Println("-=========================")
			os.Exit(0)
		}

		err = a.storage.Bb.Model(&models.Order{}).Where("symbol = ? and id = ?", req.Symbol, req.Position).Updates(&models.Order{
			ClosePrice:    price,
			CloseTime:     tick2.Timestamp,
			CreateTimeStr: time.Unix(tick2.Timestamp, 0).Format("2006-01-02 15:04:05"),
			Profit:        profit,
			Comment:       req.Comment,
		}).Error
		if err != nil {
			panic(err)
		}

		a.storage.Bb.Model(&models.Account{}).Where("account = ?", req.Account).
			UpdateColumn("balance", gorm.Expr("balance + ?", profit))
	}

	c.JSON(200, gin.H{
		"error": "",
	})
}

func (a *ApiServer) positionsTotal(c *gin.Context) {
	var req models.CloseAllReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	var count int64
	err := a.storage.Bb.Model(&models.Order{}).Where("account = ?", req.Account).
		Where("close_time = 0").Count(&count).Error
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

	if req.Ticket == 0 {
		err := a.storage.Bb.Model(&models.Order{}).
			Where("close_time = 0").
			Where("account = ?", req.Account).
			Order("create_time").Find(&orders).Error
		if err != nil {
			panic(err)
		}
	} else {
		err := a.storage.Bb.Model(&models.Order{}).
			Where("close_time = 0").
			Where("account = ?", req.Account).
			Where("id = ?", req.Ticket).
			Order("create_time").Find(&orders).Error
		if err != nil {
			panic(err)
		}
	}

	tick2 := a.storage.GetTick2()

	var items = make([]models.RespOrderPosition, 0)
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
			Symbol:       order.Symbol,
			Comment:      order.Comment,
		})
	}

	a.storage.Record(items)

	c.JSON(200, models.RespOrderPositionsGet{
		Items: items,
	})
}

func (a *ApiServer) accountInfo(c *gin.Context) {
	var req models.ReqAccountInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Fatalln(err)
	}

	account := a.storage.GetAccount(req.Account)
	c.JSON(200, account)
}

func (a *ApiServer) closeAll(c *gin.Context) {
	var req models.CloseAllReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	// 获取所有订单计算利润
	var orders []models.Order
	a.storage.Bb.Model(&models.Order{}).
		Where("account = ?", req.Account).Where("close_time = 0").Find(&orders)

	tick2 := a.storage.GetTick2()
	var myProfile float64
	for _, order := range orders {
		var profit float64
		var price float64
		if order.Type == 1 { // 如果当前订单为sell
			price = tick2.Ask // 这做空平仓
			profit = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
		} else {
			price = tick2.Bid // 做多
			profit = math.Round(((price-order.Price)*order.Volume*100000)*100) / 100
		}

		myProfile += profit
		// close
		a.storage.Bb.Model(&models.Order{}).Where("id = ?", order.ID).
			Updates(&models.Order{
				ClosePrice:   price,
				CloseTime:    tick2.Timestamp,
				CloseTimeStr: time.Unix(tick2.Timestamp, 0).Format("2006-01-02 15:04:05"),
				Profit:       profit,
				Comment:      req.Comment,
			})
	}

	a.storage.Bb.Model(&models.Account{}).Where("account = ?", req.Account).
		UpdateColumn("balance", gorm.Expr("balance + ?", myProfile))

	a.storage.RecordUp(tick2.Timestamp)
}

func (a *ApiServer) web(c *gin.Context) {
	account := c.Query("account")
	var ac models.Account
	err := a.storage.Bb.Model(&models.Account{}).Where("account = ?", account).First(&ac).Error
	if err != nil {
		c.JSON(200, gin.H{
			"error": "账户不存在",
		})
		return
	}

	getAccount := a.storage.GetAccount(account)

	c.Writer.Header().Set("Content-Type", "text/html")

	hp := html
	hp = strings.ReplaceAll(hp, "{user}", account)
	hp = strings.ReplaceAll(hp, "{rj}", fmt.Sprintf("%.2f", getAccount.InitialAmount))
	hp = strings.ReplaceAll(hp, "{balance}", fmt.Sprintf("%.2f", getAccount.Balance))
	hp = strings.ReplaceAll(hp, "{lever}", fmt.Sprintf("%d", getAccount.Lever))
	hp = strings.ReplaceAll(hp, "{margin}", fmt.Sprintf("%.2f", getAccount.Margin))
	hp = strings.ReplaceAll(hp, "{profit}", fmt.Sprintf("%.2f", getAccount.Profit))
	hp = strings.ReplaceAll(hp, "{cs}", fmt.Sprintf("%d", getAccount.LargestPosition))
	hp = strings.ReplaceAll(hp, "{ks}", fmt.Sprintf("%.2f", getAccount.LargestLoss))
	hp = strings.ReplaceAll(hp, "{yl}", fmt.Sprintf("%.2f", getAccount.LargestProfit))
	hp = strings.ReplaceAll(hp, "{tj}", fmt.Sprintf("%.2f", getAccount.Balance+getAccount.Profit-getAccount.Margin))
	hp = strings.ReplaceAll(hp, "{tj2}", fmt.Sprintf("%.2f", getAccount.Balance+getAccount.Profit))
	hp = strings.ReplaceAll(hp, "{dtmax}", fmt.Sprintf("%.2f", getAccount.FundingDynamicsMax))

	var orders []models.Order
	err = a.storage.Bb.Model(&models.Order{}).Where("close_time = 0").Where("account = ?", account).Order("create_time").Find(&orders).Error
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
			Symbol:       order.Symbol,
		})
	}

	tables := ""
	for _, item := range items {
		stime := time.Unix(item.Time, 0).Format("2006-01-02 15:04:05")
		tp := "买"
		if item.Type == 0 {
			tp = "买"
		} else {
			tp = "卖"
		}
		tables += fmt.Sprintf(`
   <tr>
		<td>%d</td>
		<td>%s</td>
		<td>%.2f</td>
		<td>%s</td>
		<td>%s</td>
		<td>%.5f</td>
		<td>%.5f</td>
		<td>%.2f</td>
	</tr>
`, item.Ticket, stime, item.Volume, tp, item.Symbol, item.PriceOpen, item.PriceCurrent, item.Profit)
	}

	hp = strings.ReplaceAll(hp, "{table}", tables)

	c.Writer.Write([]byte(hp))
}

var html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>用户订单展示</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>

<div class="container mt-5">
    <!-- 用户信息 -->
    <div class="card mb-4">
        <div class="card-header">
            用户信息
        </div>
        <div class="card-body">
            <div class="row">
                <div class="col-md-2"><strong>账户名:</strong> {user} </div>
                <div class="col-md-2"><strong>初始入金:</strong> {rj}</div>
                <div class="col-md-2"><strong>余额:</strong> {balance}</div>
                <div class="col-md-2"><strong>保证金:</strong> {margin}</div>
                <div class="col-md-2"><strong>利润:</strong> {profit}</div>
                <div class="col-md-2"><strong>杠杆:</strong> {lever}</div>
                <div class="col-md-2"><strong>最大层数:</strong> {cs}</div>
                <div class="col-md-2"><strong>最大亏损:</strong> {ks}</div>
                <div class="col-md-2"><strong>最大盈利:</strong> {yl}</div>
                <div class="col-md-2"><strong>动态金额:</strong> {tj}</div>
                <div class="col-md-2"><strong>动态金额2:</strong> {tj2}</div>
                <div class="col-md-2"><strong>动态最低:</strong> {dtmax}</div>
            </div>
        </div>
    </div>

    <!-- 当前订单 -->
    <div class="card">
        <div class="card-header">
            当前订单
        </div>
        <div class="card-body">
            <table class="table table-striped">
                <thead>
                <tr>
                    <th scope="col">订单号</th>
                    <th scope="col">时间</th>
                    <th scope="col">交易量</th>
                    <th scope="col">买/卖</th>
                    <th scope="col">货币对</th>
                    <th scope="col">买价格</th>
                    <th scope="col">当前价格</th>
                    <th scope="col">利润</th>
                </tr>
                </thead>
                <tbody>
                {table}
             
                
                <!-- 添加更多订单行 -->
                </tbody>
            </table>
        </div>
    </div>
</div>

<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
</body>
</html>
`
