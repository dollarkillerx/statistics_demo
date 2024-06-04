package models

import "gorm.io/gorm"

type Tick struct {
	ID        uint    `gorm:"primarykey"`
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
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
