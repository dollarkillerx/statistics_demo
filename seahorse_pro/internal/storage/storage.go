package storage

import (
	"fmt"
	"math"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"seahorse/internal/conf"
	"seahorse/internal/models"
)

type Storage struct {
	Rc             *conf.Conf
	Bb             *gorm.DB
	tickData       []models.Tick
	tick           models.Tick
	tickChannel    chan models.Tick
	getTickChannel chan struct{}

	mu sync.Mutex
}

func (s *Storage) Next() models.Tick {
	s.getTickChannel <- struct{}{}
	tick := <-s.tickChannel
	return tick
}

func (s *Storage) GetTick() models.Tick {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.tick
}

func (s *Storage) setTick(tick models.Tick) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tick = tick
}

func (s *Storage) heartbeat() {

}

func New(rc *conf.Conf) *Storage {
	dbConf := rc.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo", dbConf.Host,
		dbConf.User, dbConf.Password, dbConf.DbName, dbConf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Order{}, &models.Account{}, &models.OrderHistory{}, &models.OrderHistoryTick{}, &models.AccountLog{})

	rs := &Storage{Bb: db, getTickChannel: make(chan struct{}), tickChannel: make(chan models.Tick), Rc: rc}

	go rs.heartbeat()

	return rs
}

func (s *Storage) GetAccount(account string) models.Account {
	var accountInfo models.Account
	err := s.Bb.Model(&models.Account{}).Where("account = ?", account).First(&accountInfo).Error
	if err != nil {
		panic(err)
	}

	var orders []models.Order
	err = s.Bb.Model(&models.Order{}).Where("account = ?", account).
		Where("close_time = 0").Find(&orders).Error

	tick := s.GetTick()

	var myProfile float64
	var margin float64
	for idx, order := range orders {
		price := tick.Ask
		var profile float64
		if order.Type == 1 { // 如果当前订单为sell
			price = tick.Ask // 这做空平仓
			profile = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
		} else {
			price = tick.Bid // 做多
			profile = math.Round(((price-order.Price)*order.Volume*100000)*100) / 100
		}

		orders[idx].Price = price
		orders[idx].Profit = profile

		myProfile += profile
		margin += orders[idx].Margin
	}

	accountInfo.Profit = myProfile
	accountInfo.Margin = margin

	// 最大持仓
	if accountInfo.LargestPosition < len(orders) {
		accountInfo.LargestPosition = len(orders)
	}
	// 最大亏损
	if myProfile < 0 {
		if myProfile < accountInfo.LargestLoss {
			accountInfo.LargestLoss = myProfile
		}
	}
	// 最大盈利
	if myProfile > 0 {
		if myProfile > accountInfo.LargestProfit {
			accountInfo.LargestProfit = myProfile
		}
	}

	fundingDynamics := accountInfo.Balance + accountInfo.Profit
	if accountInfo.FundingDynamicsMax == 0 {
		accountInfo.FundingDynamicsMax = fundingDynamics
	}
	if fundingDynamics < accountInfo.FundingDynamicsMax {
		accountInfo.FundingDynamicsMax = fundingDynamics
	}

	// up
	err = s.Bb.Model(&models.Account{}).Where("account = ?", account).Updates(map[string]interface{}{
		"largest_position":     accountInfo.LargestPosition,
		"largest_loss":         accountInfo.LargestLoss,
		"largest_profit":       accountInfo.LargestProfit,
		"funding_dynamics_max": accountInfo.FundingDynamicsMax,
		"funding_dynamics":     fundingDynamics,
	}).Error

	return accountInfo
}
