package resp

import (
	"github.com/dollarkillerx/backend/internal/storage"
	"github.com/dollarkillerx/backend/pkg/models"
	"github.com/rs/xid"

	"time"
)

type SubscriptionPayload struct {
	SubscriptionClientID string `json:"subscription_client_id"` // 订阅账户
	StrategyCode         string `json:"strategy_code"`          // 订阅策略code

	// 当前账户信息
	ClientID  string      `json:"client_id"` // company.account: exness.10086
	Account   Account     `json:"account"`   // 账户信息
	Positions []Positions `json:"positions"` // 持仓
	History   []Positions `json:"history"`   // 历史订单
}

func (b *SubscriptionPayload) ToPositions(clientID string, storage *storage.Storage) []models.Positions {
	var result []models.Positions

	for _, v := range b.Positions {
		id := xid.New().String()
		openingTimeSystem := storage.GenTime(clientID + "openingTimeSystem")
		result = append(result, models.Positions{
			BaseModel: models.BaseModel{
				ID:        id,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			ClientID:          clientID,
			OrderID:           v.OrderID,
			Direction:         v.Direction,
			Symbol:            v.Symbol,
			Magic:             v.Magic,
			OpenPrice:         v.OpenPrice,
			Volume:            v.Volume,
			Market:            v.Market,
			Swap:              v.Swap,
			Profit:            v.Profit,
			Common:            v.Common,
			OpeningTime:       v.OpeningTime,
			ClosingTime:       v.ClosingTime,
			CommonInternal:    "",
			OpeningTimeSystem: openingTimeSystem,
			ClosingTimeSystem: 0,
		})
	}

	return result
}

func (b *SubscriptionPayload) ToHistory(clientID string, storage *storage.Storage) []models.History {
	var result []models.History

	for _, v := range b.Positions {
		id := xid.New().String()
		var openingTimeSystem = storage.GenTime(clientID + "openingTimeSystem")
		var closingTimeSystem = storage.GenTime(clientID + "closingTimeSystem")
		result = append(result, models.History{
			BaseModel: models.BaseModel{
				ID:        id,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			ClientID:          clientID,
			OrderID:           v.OrderID,
			Direction:         v.Direction,
			Symbol:            v.Symbol,
			Magic:             v.Magic,
			OpenPrice:         v.OpenPrice,
			Volume:            v.Volume,
			Market:            v.Market,
			Swap:              v.Swap,
			Profit:            v.Profit,
			Common:            v.Common,
			OpeningTime:       v.OpeningTime,
			ClosingTime:       v.ClosingTime,
			CommonInternal:    "",
			OpeningTimeSystem: openingTimeSystem,
			ClosingTimeSystem: closingTimeSystem,
		})
	}

	return result
}

type SubscriptionResponse struct {
	ClientID string `json:"client_id"` // company.account: exness.10086   推送id

	// 订阅账户信息
	SubscriptionClientID string      `json:"subscription_client_id"` // 订阅账户
	OpenPositions        []Positions `json:"open_positions"`         // 开仓订单
	ClosePosition        []Positions `json:"close_position"`         // 关仓订单
}
