package storage

import "fmt"

type CacheKey string

const (
	CacheAccount   CacheKey = "CacheAccount"   // cache 账户信息
	CachePositions CacheKey = "CachePositions" // cache 持仓
	CacheHistory   CacheKey = "CacheHistory"   // cache 历史持仓
)

func (c CacheKey) GetKey(id string) string {
	return fmt.Sprintf("%s_%s", c, id)
}
