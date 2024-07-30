package storage

import (
	"context"
	"time"

	"github.com/dollarkillerx/backend/pkg/models"
)

func (s *Storage) UpdateAccount(clientID string, account models.Account) error {
	key := CacheAccount.GetKey(clientID)

	// 1. 直接存储 account 缓存时间 1年
	return s.cache.SetEx(context.TODO(), key, account.ToJSON(), time.Hour*24*30*12).Err()
}
