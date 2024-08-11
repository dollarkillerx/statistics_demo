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
		Period    string  `json:"period"`
		MaxProfit float64 `json:"max_profit"`
		MinProfit float64 `json:"min_profit"`
	}

	var result ProfitResult

	// 存储所有统计数据的列表
	var profitData []ProfitResult

	// 当天最高利润和最低利润
	a.storage.DB().Model(&models.TimeSeriesPosition{}).
		Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
		Where("client_id = ? AND DATE(created_at) = ?", param, today).
		Scan(&result)
	result.Period = "today"
	profitData = append(profitData, result)

	// 最近5天每天的最高利润和最低利润
	for i := 0; i < 5; i++ {
		day := now.AddDate(0, 0, -i).Format("2006-01-02")
		a.storage.DB().Model(&models.TimeSeriesPosition{}).
			Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
			Where("client_id = ? AND DATE(created_at) = ?", param, day).
			Scan(&result)
		result.Period = day
		profitData = append(profitData, result)
	}

	// 本月最高利润和最低利润
	a.storage.DB().Model(&models.TimeSeriesPosition{}).
		Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
		Where("client_id = ? AND created_at >= ?", param, firstOfMonth).
		Scan(&result)
	result.Period = "this_month"
	profitData = append(profitData, result)

	// 历史最高利润和最低利润
	a.storage.DB().Model(&models.TimeSeriesPosition{}).
		Select("MAX(profit) as max_profit, MIN(profit) as min_profit").
		Where("client_id = ?", param).
		Scan(&result)
	result.Period = "all_time"
	profitData = append(profitData, result)

	// 返回利润统计数据
	resp.Return(ctx, 200, "ok", gin.H{
		"positions": pos,
		"profits":   profitData,
	})
}

func (a *ApiServer) chartsAccount(ctx *gin.Context) {
	param := strings.TrimSpace(ctx.Param("account"))

	if param == "" {
		resp.Return(ctx, 200, "key is null", nil)
		return
	}

	var tsp []models.TimeSeriesPosition
	//a.storage.DB().Model(&models.TimeSeriesPosition{}).
	//	Where("client_id = ? ", param).
	//	Where("created_at > ?", time.Now().Add(-time.Hour*24*1)).Find(&tsp)

	a.storage.DB().Model(&models.TimeSeriesPosition{}).
		Where("client_id = ? ", param).Order("created_at desc").Limit(500).Find(&tsp)
	resp.Return(ctx, 200, "ok", tsp)
}
