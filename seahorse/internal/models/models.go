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
	Balance         float64 `json:"balance"`          // 资金
	Margin          float64 `json:"margin"`           // 保证金
	Lever           int     `json:"lever"`            // 杠杆
	LargestPosition float64 `json:"largest_position"` // 最大持仓
	LargestLoss     float64 `json:"largest_loss"`     // 最大亏损
	LargestProfit   float64 `json:"largest_profit"`   // 最大盈利

	Profit float64 `json:"profit"` // 利润
}
