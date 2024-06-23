package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Symbol        string  `json:"symbol"` // 货币
	Type          int     `json:"type"`   // 0 buy，1 sell
	Price         float64 `json:"price"`
	ClosePrice    float64 `json:"close_price"`
	Volume        float64 `json:"volume"`
	CreateTime    int64   `json:"create_time"`
	CreateTimeStr string  `json:"create_time_str"`
	CloseTime     int64   `json:"close_time"`
	CloseTimeStr  string  `json:"close_time_str"`
	Profit        float64 `json:"profit"` // 利润
	Tp            float64 `json:"tp"`     // 止盈
	Sl            float64 `json:"sl"`     // 止损

	Margin float64 `json:"margin"` // 保证金

	Account string `json:"account"`

	Comment string `json:"comment"` // 备注
	Auto    bool   `json:"auto"`
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

	FundingDynamicsMax float64 `json:"funding_dynamics_max"` // 动态资金max
	FundingDynamics    float64 `json:"funding_dynamics"`     // 动态资金max

	Profit float64 `json:"profit"` // 利润
	Margin float64 `json:"margin"` // 保证金
}

type AccountLog struct {
	gorm.Model
	Account         string  `json:"account"`
	FundingDynamics float64 `json:"funding_dynamics"` // 动态资金max
}

type OrderHistory struct {
	gorm.Model
	Account      string  `json:"account"`
	CloseTime    int64   `json:"close_time"`
	CloseTimeStr string  `json:"close_time_str"`
	Profit       float64 `json:"profit"`
	Position     int     `json:"position"` // 仓位数量
	Volume       float64 `json:"volume"`   // 交易量
	Comment      string  `json:"comment"`  // 备注
}

type OrderHistoryTick struct {
	gorm.Model
	Account  string  `json:"account"`
	Time     int64   `json:"time"`
	TimeStr  string  `json:"time_str"`
	Profit   float64 `json:"profit"`   // 最值
	Position int     `json:"position"` // 仓位数量
	Volume   float64 `json:"volume"`   // 交易量
	Comment  string  `json:"comment"`  // 备注
}

type Tick struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
}

type TickItem struct {
	TimeStr string  `json:"time_str"`
	Time    int64   `json:"time"`
	Open    float64 `json:"open"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Close   float64 `json:"close"`
	Volume  float64 `json:"volume"`
	Spread  int64   `json:"spread"`
}
