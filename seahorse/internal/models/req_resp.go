package models

type ReqInit struct {
	Account string  `json:"account"`
	Balance float64 `json:"balance"` // 资金
	Lever   int     `json:"lever"`   // 杠杆
}

type RespInit struct {
	Account string `json:"account"`
	Error   string `json:"error"`
}

type ReqSymbolInfoTick struct {
	Symbol string `json:"symbol"`
}

type RespSymbolInfoTick struct {
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
	Timestamp int64   `json:"timestamp"`
}

type ReqOrderSend struct {
	Symbol   string  `json:"symbol"`
	Volume   float64 `json:"volume"`
	Type     int     `json:"type"`
	Position int     `json:"position"`
	Price    float64 `json:"price"`
}
