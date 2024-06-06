package models

import "gorm.io/gorm"

type Tick struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
	Over      bool    `json:"over"`
}

type Order struct {
	gorm.Model
	Symbol     string  `json:"symbol"` // 货币
	Type       int     `json:"type"`   // 0 buy，1 sell
	Price      float64 `json:"price"`
	ClosePrice float64 `json:"close_price"`
	Volume     float64 `json:"volume"`
	CreateTime int64   `json:"create_time"`
	CloseTime  int64   `json:"close_time"`
	Profit     float64 `json:"profit"` // 利润

	Margin float64 `json:"margin"` // 保证金

	Account string `json:"account"`
}

type Account struct {
	gorm.Model
	Account         string  `json:"account"`
	InitialAmount   float64 `json:"initial_amount"`   // 初始资金
	Balance         float64 `json:"balance"`          // 资金
	Lever           int     `json:"lever"`            // 杠杆
	LargestPosition int     `json:"largest_position"` // 最大持仓
	LargestLoss     float64 `json:"largest_loss"`     // 最大亏损
	LargestProfit   float64 `json:"largest_profit"`   // 最大盈利

	Profit float64 `json:"profit"` // 利润
	Margin float64 `json:"margin"` // 保证金
}

type OrderHistory struct {
	gorm.Model
	CloseTime int64   `json:"close_time"`
	Profit    float64 `json:"profit"`
	Position  int     `json:"position"` // 仓位数量
	Volume    float64 `json:"volume"`   // 交易量
}

type OrderHistoryTick struct {
	gorm.Model
	Time     int64   `json:"time"`
	Profit   float64 `json:"profit"`
	Position int     `json:"position"` // 仓位数量
	Volume   float64 `json:"volume"`   // 交易量
}
