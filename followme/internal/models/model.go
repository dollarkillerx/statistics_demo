package models

type Order struct {
	ID          string  `json:"id"`
	Type        int     `json:"type"` // 0: buy, 1: sell
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
	Comment     string  `json:"comment"`
	CreatedTime int     `json:"created_time"`
}
