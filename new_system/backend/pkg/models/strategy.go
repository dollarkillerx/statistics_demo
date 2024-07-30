package models

import "github.com/dollarkillerx/backend/pkg/enum"

// Strategy 策略
type Strategy struct {
	BaseModel
	ClientID        string               `json:"client_id" gorm:"column:client_id;type:varchar(255);not null"`               // company.account: exness.10086
	FollowDirection enum.FollowDirection `json:"follow_direction" gorm:"column:follow_direction;type:varchar(255);not null"` // 方向
	Proportion      float64              `json:"proportion" gorm:"column:proportion;type:decimal(10,2);not null"`            // 倍数
	ExitMode        string               `json:"exit_mode" gorm:"column:exit_mode;type:varchar(255);not null"`               // 退出模式
	Payload         string               `json:"payload" gorm:"column:payload;type:text;not null"`                           // 其他参数
}

func (Strategy) TableName() string {
	return "strategy"
}
