package models

// Statistics 统计
type Statistics struct {
	BaseModel
	Account  int64   `json:"account" gorm:"column:account;type:bigint;not null"`        // 账户
	Leverage int64   `json:"leverage" gorm:"column:leverage;type:int;not null"`         // 杠杆
	Server   string  `json:"server" gorm:"column:server;type:varchar(255);not null"`    // 服务器
	Company  string  `json:"company" gorm:"column:company;type:varchar(255);not null"`  // 公司
	Balance  float64 `json:"balance" gorm:"column:balance;type:decimal(20,8);not null"` // 余额
	Profit   float64 `json:"profit" gorm:"column:profit;type:decimal(20,8);not null"`   // 利润
	Margin   float64 `json:"margin" gorm:"column:margin;type:decimal(20,8);not null"`   // 预付款

	Payload string `json:"payload" gorm:"column:payload;type:text;not null"` // 时序数据
}
