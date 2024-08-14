package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	"log"
	"math"
	"time"

	"github.com/dollarkillerx/backend/pkg/models"
)

func (s *Storage) TimeSeriesPosition(clientID string, account models.Account, positions []models.Positions) {
	if len(positions) == 0 {
		return
	}

	// 1. 获取当前clientID 最近的一次交易 当positions持仓相同 && 价格浮动 < 2 usd 时 不记录
	var before string
	var after string
	for _, v := range positions {
		after += fmt.Sprintf("%d", v.OrderID)
	}

	beforeTSP := s.getLastTimeSeriesPosition(clientID)
	if beforeTSP == nil || beforeTSP.Payload == "" {
		marshal, _ := json.Marshal(positions)
		err := s.db.Model(&models.TimeSeriesPosition{}).Create(&models.TimeSeriesPosition{
			BaseModel: models.BaseModel{
				ID: xid.New().String(),
			},
			ClientID: clientID,
			Account:  account.Account,
			Leverage: account.Leverage,
			Server:   account.Server,
			Company:  account.Company,
			Balance:  account.Balance,
			Profit:   account.Profit,
			Margin:   account.Margin,
			Payload:  string(marshal),
		}).Error

		if err != nil {
			log.Println(err)
			return
		}

		s.createLastTimeSeriesPosition(clientID)
		return
	}

	var pos []models.Positions
	err := json.Unmarshal([]byte(beforeTSP.Payload), &pos)
	if err != nil {
		//log.Println(beforeTSP)
		//log.Println(beforeTSP.Payload)
		//log.Println(err)
		return
	}

	for _, v := range pos {
		before += fmt.Sprintf("%d", v.OrderID)
	}

	if after == before {
		//fmt.Println("-----------------------------------------------------")
		//fmt.Println("account.Profit", account.Profit)
		//fmt.Println("beforeTSP.Profit", beforeTSP.Profit)
		//fmt.Println("-----------------------------------------------------")
		if math.Abs(float64(account.Profit)-float64(beforeTSP.Profit)) < 2 {
			return
		}
	}

	// save
	marshal, _ := json.Marshal(positions)
	err = s.db.Model(&models.TimeSeriesPosition{}).Create(&models.TimeSeriesPosition{
		BaseModel: models.BaseModel{
			ID: xid.New().String(),
		},
		ClientID: clientID,
		Account:  account.Account,
		Leverage: account.Leverage,
		Server:   account.Server,
		Company:  account.Company,
		Balance:  account.Balance,
		Profit:   account.Profit,
		Margin:   account.Margin,
		Payload:  string(marshal),
	}).Error
	if err != nil {
		log.Println(err)
		return
	}
	s.createLastTimeSeriesPosition(clientID)
	return
}

func (s *Storage) getLastTimeSeriesPosition(clientID string) *models.TimeSeriesPosition {
	result, err := s.cache.Get(context.TODO(), clientID+"TimeSeriesPosition").Result()
	if err == nil {
		var tsp models.TimeSeriesPosition
		if err := json.Unmarshal([]byte(result), &tsp); err == nil {
			return &tsp
		}
	}

	return nil
}

func (s *Storage) createLastTimeSeriesPosition(clientID string) {
	var beforeTSP models.TimeSeriesPosition
	err := s.db.Model(&models.TimeSeriesPosition{}).
		Where("client_id = ?", clientID).
		Order("created_at desc").Limit(1).First(&beforeTSP).Error
	if err != nil {
		return
	}
	marshal, err := json.Marshal(beforeTSP)
	if err == nil {
		s.cache.SetEx(context.TODO(), clientID+"TimeSeriesPosition", marshal, time.Hour*24*30)
	}
}
