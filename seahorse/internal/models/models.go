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
	Volume     float64 `json:"volume"`
	CreateTime int     `json:"create_time"`
	CloseTime  int     `json:"close_time"`
	Profit     float64 `json:"profit"` // 利润
}

type Account struct {
	gorm.Model
	Account         string  `json:"account"`
	Balance         float64 `json:"balance"`          // 资金
	Lever           int     `json:"lever"`            // 杠杆
	LargestPosition float64 `json:"largest_position"` // 最大持仓
	LargestLoss     float64 `json:"largest_loss"`     // 最大亏损
	LargestProfit   float64 `json:"largest_profit"`   // 最大盈利
}
