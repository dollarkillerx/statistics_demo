package models

// Statistics 统计
type Statistics struct {
	BaseModel
	ClientID string  `json:"client_id" gorm:"column:client_id;type:varchar(255);not null"` // company.account: exness.10086
	Account  int64   `json:"account" gorm:"column:account;type:bigint;not null"`           // 账户
	Leverage int64   `json:"leverage" gorm:"column:leverage;type:int;not null"`            // 杠杆
	Server   string  `json:"server" gorm:"column:server;type:varchar(255);not null"`       // 服务器
	Company  string  `json:"company" gorm:"column:company;type:varchar(255);not null"`     // 公司
	Balance  float64 `json:"balance" gorm:"column:balance;type:decimal(20,8);not null"`    // 余额
	Margin   float64 `json:"margin" gorm:"column:margin;type:decimal(20,8);not null"`      // 预付款

	Profit    float64 `json:"profit" gorm:"column:profit;type:decimal(20,8);not null"`         // 利润
	MaxProfit float64 `json:"max_profit" gorm:"column:max_profit;type:decimal(20,8);not null"` // 最大利润
	MinProfit float64 `json:"min_profit" gorm:"column:min_profit;type:decimal(20,8);not null"` // 最小利润
}
