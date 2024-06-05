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
	Time      int64   `json:"time"`
}

type ReqOrderSend struct {
	Position int `json:"position"` // id

	Symbol string  `json:"symbol"`
	Volume float64 `json:"volume"`
	Type   int     `json:"type"`
	Price  float64 `json:"price"`

	Account string `json:"account"`
}

type ReqOrderPositionsGet struct {
	Symbol string `json:"symbol"`
	Ticket int    `json:"ticket"`

	Account string `json:"account"`
}

type RespOrderPositionsGet struct {
	Items []RespOrderPosition `json:"items"`
}

type RespOrderPosition struct {
	Ticket       uint    `json:"ticket"`
	Time         int64   `json:"time"`
	Type         int     `json:"type"`
	Volume       float64 `json:"volume"`
	PriceOpen    float64 `json:"price_open"`
	PriceCurrent float64 `json:"price_current"` // 当前价格
	Profit       float64 `json:"profit"`
	Symbol       string  `json:"symbol"` // 货币
}

type ReqAccountInfo struct {
	Account string `json:"account"`
}

type RespAccountInfo struct {
	Profit float64 `json:"profit"`
}
