package storage

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dollarkillerx/backend/pkg/models"
)

func (s *Storage) UpdateAccount(clientID string, account models.Account) error {
	key := CacheAccount.GetKey(clientID)

	// 1. 直接存储 account 缓存时间 1年
	return s.cache.SetEx(context.TODO(), key, account.ToJSON(), time.Hour*24*30*12).Err()
}

func (s *Storage) GetAccounts() []models.Account {
	var positionKey []string
	// 前缀搜索
	prefix := CacheAccount.GetKey("")
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = s.cache.Scan(context.TODO(), cursor, prefix+"*", 10).Result()
		if err != nil {
			panic(err)
		}
		positionKey = append(positionKey, keys...)
		n += len(keys)
		if cursor == 0 {
			break
		}
	}

	var result []models.Account
	for _, v := range positionKey {
		sr, err := s.cache.Get(context.TODO(), v).Result()
		if err == nil {
			var ac models.Account
			err := json.Unmarshal([]byte(sr), &ac)
			if err == nil {
				result = append(result, ac)
			}
		}
	}

	return result
}
