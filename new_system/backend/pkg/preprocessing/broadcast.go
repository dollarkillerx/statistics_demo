package preprocessing

import (
	"fmt"
	"github.com/dollarkillerx/backend/internal/storage"
	"github.com/dollarkillerx/backend/pkg/enum"
	"github.com/dollarkillerx/backend/pkg/models"
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/rs/xid"
	"sort"
	"time"
)

func AccountToModel(clientID string, a resp.Account) models.Account {
	id := xid.New().String()
	return models.Account{
		BaseModel: models.BaseModel{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ClientID: clientID,
		Account:  a.Account,
		Leverage: a.Leverage,
		Server:   a.Server,
		Company:  a.Company,
		Balance:  a.Balance,
		Profit:   a.Profit,
		Margin:   a.Margin,
	}
}

func BroadcastPayloadToPositions(clientID string, storage *storage.Storage, b *resp.BroadcastPayload) []models.Positions {
	var result []models.Positions

	for _, v := range b.Positions {
		id := xid.New().String()
		openingTimeSystem := storage.GenTime(fmt.Sprintf("%s_%s_%d", clientID, "openingTimeSystem", v.OrderID))
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

func BroadcastPayloadToHistory(clientID string, storage *storage.Storage, b *resp.BroadcastPayload) []models.History {
	var result []models.History

	for _, v := range b.Positions {
		id := xid.New().String()
		var openingTimeSystem = storage.GenTime(fmt.Sprintf("%s_%s_%d", clientID, "openingTimeSystem", v.OrderID))
		var closingTimeSystem = storage.GenTime(fmt.Sprintf("%s_%s_%d", clientID, "closingTimeSystem", v.OrderID))
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

func SubscriptionPayloadToPositions(clientID string, storage *storage.Storage, b *resp.SubscriptionPayload) []models.Positions {
	var result []models.Positions

	for _, v := range b.Positions {
		id := xid.New().String()
		openingTimeSystem := storage.GenTime(fmt.Sprintf("%s_%s_%d", clientID, "openingTimeSystem", v.OrderID))
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

func SubscriptionPayloadToHistory(clientID string, storage *storage.Storage, b *resp.SubscriptionPayload) []models.History {
	var result []models.History

	for _, v := range b.Positions {
		id := xid.New().String()
		var openingTimeSystem = storage.GenTime(fmt.Sprintf("%s_%s_%d", clientID, "openingTimeSystem", v.OrderID))
		var closingTimeSystem = storage.GenTime(fmt.Sprintf("%s_%s_%d", clientID, "closingTimeSystem", v.OrderID))
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

func HistoryToHistory(history []resp.History) ([]models.History, error) {
	// 将历史记录按 PositionId 分组
	mp := make(map[int][]resp.History)
	for _, v := range history {
		mp[v.PositionId] = append(mp[v.PositionId], v)
	}

	var res []models.History

	// 遍历每个 PositionId 的历史记录
	for _, histories := range mp {
		if len(histories) != 2 {
			// 如果不是平仓和开仓记录，跳过
			if histories[0].Ticket == histories[0].PositionId {
				continue
			}
			// 根据开仓和平仓记录创建 Positions
			position := models.History{
				BaseModel: models.BaseModel{
					ID:        xid.New().String(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				OrderID:           int64(histories[0].PositionId),
				Direction:         enum.Direction(histories[0].Type),
				Symbol:            histories[0].Symbol,
				Magic:             int64(histories[0].Magic),
				OpenPrice:         histories[0].PriceCurrent,
				Volume:            histories[0].VolumeInitial,
				Market:            histories[0].PriceCurrent,
				Swap:              0,                                                     // 库存费可能需要额外计算
				Profit:            histories[0].PriceCurrent - histories[0].PriceCurrent, // 假设利润是平仓价格减去开仓价格
				Common:            histories[0].Comment,
				OpeningTime:       int64(histories[0].TimeSetup),
				ClosingTime:       int64(histories[0].TimeSetup),
				CommonInternal:    "", // 需要根据实际情况设置
				ClosingTimeSystem: time.Now().Unix(),
			}

			res = append(res, position)
			continue
		}

		// 按时间排序
		sort.Slice(histories, func(i, j int) bool {
			return histories[i].TimeSetup < histories[j].TimeSetup
		})

		opening := histories[0]
		closing := histories[1]

		// 根据开仓和平仓记录创建 Positions
		position := models.History{
			BaseModel: models.BaseModel{
				ID:        xid.New().String(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			OrderID:           int64(opening.PositionId),
			Direction:         enum.Direction(opening.Type),
			Symbol:            opening.Symbol,
			Magic:             int64(opening.Magic),
			OpenPrice:         opening.PriceCurrent,
			Volume:            opening.VolumeInitial,
			Market:            closing.PriceCurrent,
			Swap:              0,                                           // 库存费可能需要额外计算
			Profit:            closing.PriceCurrent - opening.PriceCurrent, // 假设利润是平仓价格减去开仓价格
			Common:            closing.Comment,
			OpeningTime:       int64(opening.TimeSetup),
			ClosingTime:       int64(closing.TimeSetup),
			CommonInternal:    "", // 需要根据实际情况设置
			ClosingTimeSystem: time.Now().Unix(),
		}

		res = append(res, position)
	}

	return res, nil
}
