package internal

type PostData struct {
	Ticker    string `json:"ticker"`
	Action    string `json:"action"`
	Contracts string `json:"contracts"`
	Price     string `json:"price"`
}
