package storage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

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

	s.mu.Lock()
	defer s.mu.Unlock()
	s.tick = tick

	delRint()
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

var mu sync.Mutex
var rint int

func addRint() {
	mu.Lock()
	defer mu.Unlock()
	rint++
}

func delRint() {
	mu.Lock()
	defer mu.Unlock()
	rint--
}

func next() bool {
	mu.Lock()
	defer mu.Unlock()

	if rint <= 0 {
		return true
	}

	return false
}

func (s *Storage) heartbeat() {
	open, err := os.Open(s.Rc.CsvPath)
	if err != nil {
		panic(err)
	}
	defer open.Close()

	reader := bufio.NewReader(open)

loop:
	for {
		select {
		case <-s.getTickChannel:

			if !next() {
				continue
			}

			line, _, err := reader.ReadLine()
			if err != nil {
				break loop
			}

			lines := string(bytes.TrimSpace(line))
			var tickItem models.TickItem
			err = json.Unmarshal([]byte(lines), &tickItem)
			if err == nil {
				s.tickChannel <- models.Tick{
					Symbol:    "EURUSD",
					Timestamp: tickItem.Time,
					Ask:       tickItem.Open + GenerateRandomNumber(), // 定义点差
					Bid:       tickItem.Open,
				}
				addRint()
				s.tickChannel <- models.Tick{
					Symbol:    "EURUSD",
					Timestamp: tickItem.Time + 10,
					Ask:       tickItem.Close + GenerateRandomNumber(),
					Bid:       tickItem.Close,
				}
				addRint()
			}
		}
	}
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

// 随机点查 70% 点差为10点
func GenerateRandomNumber() float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成0到99之间的随机数
	randomValue := rand.Intn(100)

	if randomValue < 70 {
		return 0.0001
	} else {
		// 30%的时间生成20或30
		if rand.Intn(2) == 0 {
			return 0.0002
		} else {
			return 0.0003
		}
	}
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

func (s *Storage) JxGd() {
	var orders []models.Order
	err := s.Bb.Model(&models.Order{}).
		Where("close_time = 0").Find(&orders).Error
	if err != nil {
		log.Println(err)
		return
	}

	if len(orders) == 0 {
		return
	}

	var jxOrder []models.Order
	for _, order := range orders {
		if order.Sl == 0 && order.Tp == 0 {
			continue
		}

		if order.Sl != 0 {
			jxOrder = append(jxOrder, order)
		}

		if order.Tp != 0 {
			jxOrder = append(jxOrder, order)
		}
	}

	if len(jxOrder) == 0 {
		return
	}

	for i, v := range jxOrder {
		if v.Sl != 0 {
			if v.Type == 0 {
				if s.GetTick().Bid <= v.Sl {
					s.closeOrder(jxOrder[i])
				}
			} else {
				if s.GetTick().Ask >= v.Sl {
					s.closeOrder(jxOrder[i])
				}
			}
		}

		if v.Tp != 0 {
			if v.Type == 0 {
				if s.GetTick().Bid >= v.Tp {
					s.closeOrder(jxOrder[i])
				}
			} else {
				if s.GetTick().Ask <= v.Tp {
					s.closeOrder(jxOrder[i])
				}
			}
		}
	}
}

func (s *Storage) closeOrder(order models.Order) {
	tick := s.GetTick()

	var price float64
	var profit float64
	if order.Type == 1 { // 如果当前订单为sell
		price = tick.Ask // 这做空平仓
		profit = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
	} else {
		price = tick.Bid // 做多
		profit = math.Round(((price-order.Price)*order.Volume*100000)*100) / 100
	}

	err := s.Bb.Model(&models.Order{}).Where("id = ?", order.ID).Updates(map[string]interface{}{
		"close_price":    price,
		"close_time":     tick.Timestamp,
		"close_time_str": time.Unix(tick.Timestamp, 0).Format("2006-01-02 15:04:05"),
		"profit":         profit,
		"comment":        "tp/sl",
		"auto":           true,
	}).Error
	if err != nil {
		panic(err)
	}
}

// ProfitAndLossCalculation 计算盈亏
func (s *Storage) ProfitAndLossCalculation(account string) {
	s.Bb.Transaction(func(tx *gorm.DB) error {
		var accountInfo models.Account
		err := tx.Model(&models.Account{}).Where("account = ?", account).First(&accountInfo).Error
		if err != nil {
			panic(err)
		}

		var orders []models.Order
		err = tx.Model(&models.Order{}).Where("account = ?", account).
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

		// 净值日志
		alog := models.AccountLog{
			Account:         account,
			OrderTotal:      len(orders),
			FundingDynamics: accountInfo.Balance + accountInfo.Profit,
		}

		s.Bb.Create(&alog)

		return nil
	})
}
