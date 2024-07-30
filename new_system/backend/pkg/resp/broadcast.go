package resp

import (
	"github.com/dollarkillerx/backend/pkg/enum"
	"github.com/dollarkillerx/backend/pkg/models"
	"github.com/rs/xid"
	"time"
)

type BroadcastPayload struct {
	ClientID  string      `json:"client_id"` // company.account: exness.10086
	Account   Account     `json:"account"`   // 账户信息
	Positions []Positions `json:"positions"` // 持仓
	History   []Positions `json:"history"`   // 历史订单
}

// Account 账户
type Account struct {
	Account  int64   `json:"account"`  // 账户
	Leverage int64   `json:"leverage"` // 杠杠
	Server   string  `json:"server"`   // 服务器
	Company  string  `json:"company"`  // company
	Balance  float64 `json:"balance"`  // 余额
	Profit   float64 `json:"profit"`   // 利润
	Margin   float64 `json:"margin"`   // 预付款
}

func (a *Account) ToModel(clientID string) models.Account {
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

// Positions 持仓
type Positions struct {
	OrderID     int64          `json:"order_id"`     // 持仓ID
	Direction   enum.Direction `json:"direction"`    // 方向
	Symbol      string         `json:"symbol"`       // 币种
	Magic       int64          `json:"magic"`        // 魔术手
	OpenPrice   float64        `json:"open_price"`   // 开仓价格
	Volume      float64        `json:"volume"`       // 数量
	Market      float64        `json:"market"`       // 市价
	Swap        float64        `json:"swap"`         // 库存费
	Profit      float64        `json:"profit"`       // 利润
	Common      string         `json:"common"`       // 注释
	OpeningTime int64          `json:"opening_time"` // 开仓时间市商
	ClosingTime int64          `json:"closing_time"` // 平仓时间市商

	CommonInternal    string `json:"common_internal"`     // 系统内部注释
	OpeningTimeSystem int64  `json:"opening_time_system"` // 开仓时间系统
	ClosingTimeSystem int64  `json:"closing_time_system"` // 平仓时间系统
}

func (b *BroadcastPayload) ToPositions(clientID string, oldPositions []models.Positions) []models.Positions {
	var result []models.Positions

	for _, v := range b.Positions {
		id := xid.New().String()
		var openingTimeSystem int64
		pos := models.GetPositionsByID(clientID, v.OrderID, oldPositions)
		if pos != nil {
			id = pos.ID
			openingTimeSystem = pos.OpeningTimeSystem
		} else {
			openingTimeSystem = time.Now().Unix()
		}
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

func (b *BroadcastPayload) ToHistory(clientID string, oldHistory []models.History) []models.History {
	var result []models.History

	for _, v := range b.Positions {
		id := xid.New().String()
		var openingTimeSystem int64
		pos := models.GetPositionsByID(clientID, v.OrderID, oldPositions)
		if pos != nil {
			id = pos.ID
			openingTimeSystem = pos.OpeningTimeSystem
		} else {
			openingTimeSystem = time.Now().Unix()
		}
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
