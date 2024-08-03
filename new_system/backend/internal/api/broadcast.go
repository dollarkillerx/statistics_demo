package api

import (
	"github.com/dollarkillerx/backend/pkg/preprocessing"
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (a *ApiServer) broadcast(ctx *gin.Context) {
	var input resp.BroadcastPayload
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
	positions := preprocessing.BroadcastPayloadToPositions(input.ClientID, a.storage, &input)
	a.storage.UpdatePositions(input.ClientID, positions)

	// 3. 更新历史订单
	history, _ := preprocessing.HistoryToHistory(input.History)
	a.storage.UpdateHistory(input.ClientID, history)

	// log
	go a.storage.TimeSeriesPosition(input.ClientID, preprocessing.AccountToModel(input.ClientID, input.Account), positions)

	resp.Return(ctx, 200, "ok", nil)
}
