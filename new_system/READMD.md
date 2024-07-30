# APi

### 存储设计

- 热数据 dragonflydb
  - 账户数据
  - 持仓数据
- 冷数据 pgsql
  - 历史数据

第一版 所有数据都会存储在pgsql 中

### 提交基础信息

- 广播 broadcast `post` 

``` 
type BroadcastPayload struct {
	ClientID  string      `json:"client_id"` // company.account: exness.10086
	Account   Account     `json:"account"`   // 账户信息
	Positions []Positions `json:"positions"` // 持仓
	History   []Positions `json:"history"`   // 历史订单
}

// Account 账户
type Account struct {
	Account  int64   `json:"account"`  // 账户
	Leverage int     `json:"leverage"` // 杠杠
	Server   string  `json:"server"`   // 服务器
	Company  string  `json:"company"`  // company
	Balance  float64 `json:"balance"`  // 余额
	Profit   float64 `json:"profit"`   // 利润
	Margin   float64 `json:"margin"`   // 预付款
}

// Positions 持仓
type Positions struct {
	OrderID     string         `json:"order_id"`     // 持仓ID
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
```

- 订阅 subscription `post`

request: 
``` 
type Subscription struct {
	SubscriptionClientID string `json:"subscription_client_id"` // 订阅账户 

	ClientID  string      `json:"client_id"` // company.account: exness.10086
	Account   Account     `json:"account"`   // 账户信息
	Positions []Positions `json:"positions"` // 持仓
	History   []Positions `json:"history"`   // 历史订单
}
```

response: 
``` 
type SubscriptionPayload struct {
	SubscriptionClientID string `json:"subscription_client_id"` // 订阅账户
	StrategyCode         string `json:"strategy_code"`          // 订阅策略code

	// 当前账户信息
	ClientID  string      `json:"client_id"` // company.account: exness.10086
	Account   Account     `json:"account"`   // 账户信息
	Positions []Positions `json:"positions"` // 持仓
	History   []Positions `json:"history"`   // 历史订单
}

type SubscriptionResponse struct {
	ClientID string `json:"client_id"` // company.account: exness.10086   推送id

	// 订阅账户信息
	SubscriptionClientID string      `json:"subscription_client_id"` // 订阅账户
	OpenPositions        []Positions `json:"open_positions"`         // 开仓订单
	ClosePosition        []Positions `json:"close_position"`         // 关仓订单
}
```

- 错误 `error`

request:
``` 
type ErrorPayload struct {
	ClientID string `json:"client_id"` // company.account: exness.10086
	ErrMsg   string `json:"err_msg"`   // error message
}
```

### 表设计

`new_system/backend/pkg/models`


