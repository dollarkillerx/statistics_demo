package api

import (
	"github.com/dollarkillerx/backend/pkg/models"
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/gin-gonic/gin"
	"time"

	"strings"
)

func (a *ApiServer) accounts(ctx *gin.Context) {
	positions := a.storage.GetAccounts()
	resp.Return(ctx, 200, "ok", positions)
}

func (a *ApiServer) account(ctx *gin.Context) {
	param := strings.TrimSpace(ctx.Param("account"))

	if param == "" {
		resp.Return(ctx, 200, "key is null", nil)
		return
	}

	// 获取基础数据
	pos := a.storage.GetPositionsByID(param)

	// 今日日期
	today := time.Now().Format("2006-01-02")

	// 最近5天日期
	now := time.Now()

	// 本月的开始日期
	firstOfMonth := now.Format("2006-01-01")

	type ProfitResult struct {
		MaxProfit float64 `json:"max_profit"`
		MinProfit float64 `json:"min_profit"`
	}

	var result ProfitResult

	// 存储所有统计数据
	profitData := make(map[string]ProfitResult)

	// 当天最高利润和最低利润
	a.storage.DB().Model(&models.TimeSeriesPosition{}).
		Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
		Where("client_id = ? AND DATE(created_at) = ?", param, today).
		Scan(&result)
	profitData["today"] = result

	// 最近5天每天的最高利润和最低利润
	for i := 0; i < 5; i++ {
		day := now.AddDate(0, 0, -i).Format("2006-01-02")
		a.storage.DB().Model(&models.TimeSeriesPosition{}).
			Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
			Where("client_id = ? AND DATE(created_at) = ?", param, day).
			Scan(&result)
		profitData[day] = result
	}

	// 本月最高利润和最低利润
	a.storage.DB().Model(&models.TimeSeriesPosition{}).
		Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
		Where("client_id = ? AND created_at >= ?", param, firstOfMonth).
		Scan(&result)
	profitData["this_month"] = result

	// 历史最高利润和最低利润
	a.storage.DB().Model(&models.TimeSeriesPosition{}).
		Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
		Where("client_id = ?", param).
		Scan(&result)
	profitData["all_time"] = result

	// 返回利润统计数据
	resp.Return(ctx, 200, "ok", gin.H{
		"positions": pos,
		"profits":   profitData,
	})
}
