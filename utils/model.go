package utils

import "time"

// ExchangeRate 汇率
type ExchangeRate struct {
	Time     time.Time // 时间
	Timeunix int64     // 时间戳
	Price    float64   // 价格
}
