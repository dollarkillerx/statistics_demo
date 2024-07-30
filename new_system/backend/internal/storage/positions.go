package storage

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/dollarkillerx/backend/pkg/models"
)

func (s *Storage) GenTime(key string) int64 {
	key = "GEN_" + key
	now := time.Now().Unix()
	result, err := s.cache.Get(context.TODO(), key).Result()

	if err == nil {
		s.cache.SetEx(context.TODO(), key, strconv.Itoa(int(now)), time.Hour*24*30)
		return now
	} else {
		s.cache.Expire(context.TODO(), key, time.Hour*24*30)
	}

	atoi, err := strconv.Atoi(result)
	if err != nil {
		s.cache.SetEx(context.TODO(), key, strconv.Itoa(int(now)), time.Hour*24*30)
		return now
	}

	return int64(atoi)
}

// UpdatePositions 更新持仓
func (s *Storage) UpdatePositions(clientID string, positions []models.Positions) {
	key := CachePositions.GetKey(clientID)
	marshal, _ := json.Marshal(positions)
	// 	1. 直接存储 positions 缓存时间 30日
	s.cache.SetEx(context.TODO(), key, string(marshal), time.Hour*24*30)
}

func (s *Storage) GetPositionsByID(id string) []models.Positions {
	result, err := s.cache.Get(context.TODO(), id).Result()
	if err != nil {
		return nil
	}

	var res []models.Positions
	err = json.Unmarshal([]byte(result), &res)
	if err != nil {
		return nil
	}

	return res
}

// UpdateHistory 更新历史
func (s *Storage) UpdateHistory(clientID string, positions []models.History) {
	key := CacheHistory.GetKey(clientID)
	marshal, _ := json.Marshal(positions)
	// 	1. 直接存储 positions 缓存时间 30日
	s.cache.SetEx(context.TODO(), key, string(marshal), time.Hour*24*30)
}
