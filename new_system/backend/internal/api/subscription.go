package api

import (
	"github.com/dollarkillerx/backend/pkg/enum"
	"github.com/dollarkillerx/backend/pkg/preprocessing"
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (a *ApiServer) subscription(ctx *gin.Context) {
	var input resp.SubscriptionPayload
	var result resp.SubscriptionResponse
	if err := ctx.ShouldBindJSON(&input); err != nil {
		resp.Return(ctx, 400, err.Error(), nil)
		return
	}

	// 1. 更新 account
	if err := a.storage.UpdateAccount(input.ClientID, preprocessing.AccountToModel(input.ClientID, input.Account)); err != nil {
		log.Error().Msgf("update account error: %s", err.Error())
		return
	}
	// 2. 更新当前持仓
	positions := preprocessing.SubscriptionPayloadToPositions(input.ClientID, a.storage, &input)
	a.storage.UpdatePositions(input.ClientID, positions)

	// 3. 更新历史订单
	history := preprocessing.SubscriptionPayloadToHistory(input.ClientID, a.storage, &input)
	a.storage.UpdateHistory(input.ClientID, history)

	a.storage.UpdateHistory(input.ClientID, history)

	// StrategyCode 内置

	// Reverse1: 简单的反向
	// Reverse2: 简单的反向 100止盈

	result.ClientID = input.ClientID
	result.SubscriptionClientID = input.SubscriptionClientID
	// 4. 获得订阅者的当前持仓订单
	pos := a.storage.GetPositionsByID(input.SubscriptionClientID)
	if len(pos) > 0 {
		for i, _ := range pos {

			if input.StrategyCode == "Reverse" {
				if pos[i].Direction == enum.BUY {
					pos[i].Direction = enum.SELL
				} else {
					pos[i].Direction = enum.BUY
				}
			}

			result.OpenPositions = append(result.OpenPositions, resp.Positions{
				OrderID:           pos[i].OrderID,
				Direction:         pos[i].Direction,
				Symbol:            pos[i].Symbol,
				Magic:             pos[i].Magic,
				OpenPrice:         pos[i].OpenPrice,
				Volume:            pos[i].Volume,
				Market:            pos[i].Market,
				Swap:              pos[i].Swap,
				Profit:            pos[i].Profit,
				Common:            pos[i].Common,
				OpeningTime:       pos[i].OpeningTime,
				ClosingTime:       pos[i].ClosingTime,
				OpeningTimeSystem: pos[i].OpeningTimeSystem,
				ClosingTimeSystem: pos[i].ClosingTimeSystem,
				CommonInternal:    pos[i].CommonInternal,
			})
		}
	}

	// log
	go a.storage.TimeSeriesPosition(input.ClientID, preprocessing.AccountToModel(input.ClientID, input.Account), positions)

	resp.Return(ctx, 200, "success", result)
}
