package api

import (
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (a *ApiServer) subscription(ctx *gin.Context) {
	var input resp.SubscriptionPayload
	if err := ctx.ShouldBindJSON(&input); err != nil {
		resp.Return(ctx, 400, err.Error(), nil)
		return
	}

	// 1. 更新 account
	if err := a.storage.UpdateAccount(input.ClientID, input.Account.ToModel(input.ClientID)); err != nil {
		log.Error().Msgf("update account error: %s", err.Error())
		return
	}
	// 2. 更新当前持仓
	positions := input.ToPositions(input.ClientID, a.storage)
	a.storage.UpdatePositions(input.ClientID, positions)

	// 3. 更新历史订单
	history := input.ToHistory(input.ClientID, a.storage)
	a.storage.UpdateHistory(input.ClientID, history)

	// StrategyCode 内置
	// Reverse1: 简单的反向
	// Reverse2: 简单的反向

}
