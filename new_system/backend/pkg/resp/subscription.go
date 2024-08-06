package resp

type SubscriptionPayload struct {
	SubscriptionClientID string `json:"subscription_client_id"` // 订阅账户
	StrategyCode         string `json:"strategy_code"`          // 订阅策略code

	// 当前账户信息
	ClientID  string      `json:"client_id"` // company.account: exness.10086
	Account   Account     `json:"account"`   // 账户信息
	Positions []Positions `json:"positions"` // 持仓
	History   []History   `json:"history"`   // 历史订单
}

type SubscriptionResponse struct {
	ClientID string `json:"client_id"` // company.account: exness.10086   推送id

	// 订阅账户信息
	SubscriptionClientID string      `json:"subscription_client_id"` // 订阅账户
	OpenPositions        []Positions `json:"open_positions"`         // 开仓订单
	ClosePosition        []Positions `json:"close_position"`         // 关仓订单
}
