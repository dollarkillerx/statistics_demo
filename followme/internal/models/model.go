package models

type Order struct {
	ID          string  `json:"id"`
	Type        int     `json:"type"` // 0: buy, 1: sell
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
	Comment     string  `json:"comment"`
	Symbol      string  `json:"symbol"`
	Magic       int     `json:"magic"`
	Sl          float64 `json:"sl"` // 订单止损价位点位
	TP          float64 `json:"tp"` // 订单止盈价位点位
	CreatedTime int     `json:"created_time"`
}
