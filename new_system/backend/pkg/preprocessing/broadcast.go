package preprocessing

import (
	"github.com/dollarkillerx/backend/internal/storage"
	"github.com/dollarkillerx/backend/pkg/models"
	"github.com/dollarkillerx/backend/pkg/resp"
	"github.com/rs/xid"
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

func BroadcastPayloadToHistory(clientID string, storage *storage.Storage, b *resp.BroadcastPayload) []models.History {
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

func SubscriptionPayloadToPositions(clientID string, storage *storage.Storage, b *resp.SubscriptionPayload) []models.Positions {
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

func SubscriptionPayloadToHistory(clientID string, storage *storage.Storage, b *resp.SubscriptionPayload) []models.History {
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
