package resp

import (
	"github.com/dollarkillerx/backend/pkg/enum"
)

type BroadcastPayload struct {
	ClientID  string      `json:"client_id"` // company.account: exness.10086
	Account   Account     `json:"account"`   // 账户信息
	Positions []Positions `json:"positions"` // 持仓
	History   []History   `json:"history"`   // 历史订单
}

// Account 账户
type Account struct {
	Account  int64   `json:"account"`  // 账户
	Leverage int64   `json:"leverage"` // 杠杠
	Server   string  `json:"server"`   // 服务器
	Company  string  `json:"company"`  // company
	Balance  float64 `json:"balance"`  // 余额
	Profit   float64 `json:"profit"`   // 利润
	Margin   float64 `json:"margin"`   // 预付款
}

// Positions 持仓
type Positions struct {
	OrderID     int64          `json:"order_id"`     // 持仓ID
	Direction   enum.Direction `json:"direction"`    // 方向
	Symbol      string         `json:"symbol"`       // 币种
	Magic       int64          `json:"magic"`        // 魔术手
	OpenPrice   float64        `json:"open_price"`   // 开仓价格
	Volume      float64        `json:"volume"`       // 数量
	Market      float64        `json:"market"`       // 市价
	Swap        float64        `json:"swap"`         // 库存费
	Profit      float64        `json:"profit"`       // 利润
	Common      string         `json:"common"`       // 注释
	OpeningTime int64          `json:"opening_time"` // 开仓时间市商
	ClosingTime int64          `json:"closing_time"` // 平仓时间市商

	CommonInternal    string `json:"common_internal"`     // 系统内部注释
	OpeningTimeSystem int64  `json:"opening_time_system"` // 开仓时间系统
	ClosingTimeSystem int64  `json:"closing_time_system"` // 平仓时间系统
}

type History struct {
	Ticket        int     `json:"ticket"`
	TimeSetup     int     `json:"time_setup"`
	Type          string  `json:"type"`
	Magic         int     `json:"magic"`
	PositionId    int     `json:"position_id"`
	VolumeInitial float64 `json:"volume_initial"`
	PriceCurrent  float64 `json:"price_current"`
	Symbol        string  `json:"symbol"`
	Comment       string  `json:"comment"`
}
