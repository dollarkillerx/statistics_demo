package models

import "github.com/dollarkillerx/backend/pkg/enum"

// Positions 持仓
type Positions struct {
	BaseModel
	ClientID       string         `json:"client_id" gorm:"column:client_id;type:varchar(255);not null"`    // company.account: exness.10086
	OrderID        int64          `json:"order_id" gorm:"column:order_id;type:type:bigint;not null"`       // 持仓ID
	Direction      enum.Direction `json:"direction" gorm:"column:direction;type:varchar(255);not null"`    // 方向
	Symbol         string         `json:"symbol" gorm:"column:symbol;type:varchar(50);not null"`           // 币种
	Magic          int64          `json:"magic" gorm:"column:magic;type:bigint;not null"`                  // 魔术手
	OpenPrice      float64        `json:"open_price" gorm:"column:open_price;type:decimal(20,8);not null"` // 开仓价格
	Volume         float64        `json:"volume" gorm:"column:volume;type:decimal(20,8);not null"`         // 数量
	Market         float64        `json:"market" gorm:"column:market;type:decimal(20,8);not null"`         // 市价
	Swap           float64        `json:"swap" gorm:"column:swap;type:decimal(20,8);not null"`             // 库存费
	Profit         float64        `json:"profit" gorm:"column:profit;type:decimal(20,8);not null"`         // 利润
	Common         string         `json:"common" gorm:"column:common;type:varchar(255);not null"`          // 注释
	OpeningTime    int64          `json:"opening_time" gorm:"column:opening_time;type:bigint;not null"`    // 开仓时间市商
	ClosingTime    int64          `json:"closing_time" gorm:"column:closing_time;type:bigint;not null"`    // 平仓时间市商
	CommonInternal string         `json:"common_internal" gorm:"column:common_internal;type:text"`         // 系统内部注释

	OpeningTimeSystem int64 `json:"opening_time_system" gorm:"column:opening_time_system;type:bigint"` // 开仓时间系统
	ClosingTimeSystem int64 `json:"closing_time_system" gorm:"column:closing_time_system;type:bigint"` // 平仓时间系统
}

// TableName 表名
func (Positions) TableName() string {
	return "positions"
}

func GetPositionsByID(clientID string, orderID int64, positions []Positions) *Positions {
	for idx, val := range positions {
		if val.OrderID == orderID && clientID == val.ClientID {
			return &positions[idx]
		}
	}

	return nil
}
