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
	// 3. 更新历史订单
	history, _ := preprocessing.HistoryToHistory(input.History)

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
				OpeningTime:       pos[i].OpeningTimeSystem,
				ClosingTime:       pos[i].ClosingTime,
				OpeningTimeSystem: pos[i].OpeningTimeSystem,
				ClosingTimeSystem: pos[i].ClosingTimeSystem,
				CommonInternal:    pos[i].CommonInternal,
			})
		}
	}

	his := a.storage.GetHistoryByID(input.SubscriptionClientID)
	if len(his) > 0 {
		for i, _ := range his {

			if input.StrategyCode == "Reverse" {
				if his[i].Direction == enum.BUY {
					his[i].Direction = enum.SELL
				} else {
					his[i].Direction = enum.BUY
				}
			}

			result.ClosePosition = append(result.ClosePosition, resp.Positions{
				OrderID:           his[i].OrderID,
				Direction:         his[i].Direction,
				Symbol:            his[i].Symbol,
				Magic:             his[i].Magic,
				OpenPrice:         his[i].OpenPrice,
				Volume:            his[i].Volume,
				Market:            his[i].Market,
				Swap:              his[i].Swap,
				Profit:            his[i].Profit,
				Common:            his[i].Common,
				OpeningTime:       his[i].OpeningTime,
				ClosingTime:       his[i].ClosingTime,
				OpeningTimeSystem: his[i].OpeningTimeSystem,
				ClosingTimeSystem: his[i].ClosingTimeSystem,
				CommonInternal:    his[i].CommonInternal,
			})
		}
	}

	// log
	go a.storage.TimeSeriesPosition(input.ClientID, preprocessing.AccountToModel(input.ClientID, input.Account), positions)

	resp.Return(ctx, 200, "success", result)
}
