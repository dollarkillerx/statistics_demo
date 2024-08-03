package test

import (
	"fmt"
	"github.com/dollarkillerx/backend/pkg/enum"
	"sort"
	"testing"
	"time"
)

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

type History struct {
	Ticket        int     `json:"ticket"`
	TimeSetup     int     `json:"time_setup"`
	Type          string  `json:"type"`
	Magic         int     `json:"magic"`
	PositionId    int     `json:"position_id"`
	VolumeInitial float64 `json:"volume_initial"`
	PriceCurrent  float64 `json:"price_current"`
	Symbol        string  `json:"symbol"`
	Comment       string  `json:"comment"`
}

func HistoryToPositions(history []History) ([]Positions, error) {
	// 将历史记录按 PositionId 分组
	mp := make(map[int][]History)
	for _, v := range history {
		mp[v.PositionId] = append(mp[v.PositionId], v)
	}

	var res []Positions

	// 遍历每个 PositionId 的历史记录
	for _, histories := range mp {
		if len(histories) != 2 {
			// 如果不是平仓和开仓记录，跳过
			if histories[0].Ticket == histories[0].PositionId {
				continue
			}
			// 根据开仓和平仓记录创建 Positions
			position := Positions{
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
		position := Positions{
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

func TestOrder(t *testing.T) {
	history := []History{
		{Ticket: 257615429, TimeSetup: 1722642658, Type: "BUY", Magic: 0, PositionId: 257615429, VolumeInitial: 0.01, PriceCurrent: 61551.73, Symbol: "BTCUSDm", Comment: ""},
		{Ticket: 257615433, TimeSetup: 1722642671, Type: "SELL", Magic: 0, PositionId: 257615429, VolumeInitial: 0.01, PriceCurrent: 61528.31, Symbol: "BTCUSDm", Comment: ""},
		{Ticket: 257612429, TimeSetup: 1722642658, Type: "BUY", Magic: 0, PositionId: 257615419, VolumeInitial: 0.01, PriceCurrent: 61551.73, Symbol: "BTCUSDm", Comment: ""},
		{Ticket: 257645433, TimeSetup: 1722642671, Type: "SELL", Magic: 0, PositionId: 257615419, VolumeInitial: 0.01, PriceCurrent: 61528.31, Symbol: "BTCUSDm", Comment: ""},
		{Ticket: 2576154239, TimeSetup: 1722642658, Type: "BUY", Magic: 0, PositionId: 2576154239, VolumeInitial: 0.01, PriceCurrent: 61551.73, Symbol: "BTCUSDm", Comment: ""},
	}

	fmt.Println(history)
	positions, err := HistoryToPositions(history)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, position := range positions {
			fmt.Println(position)
		}
	}
}
