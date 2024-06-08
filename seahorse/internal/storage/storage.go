package storage

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
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

func (s *Storage) GetTick() models.Tick {
	s.getTickChannel <- struct{}{}
	tick := <-s.tickChannel
	return tick
	//s.mu.Lock()
	//defer s.mu.Unlock()
	//return s.tick
}

func (s *Storage) GetTick2() models.Tick {
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
	open, err := os.Open(s.Rc.CsvPath)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	reader := csv.NewReader(open)
	record, err := reader.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println(record)

	go func() {
		s.getTickChannel <- struct{}{}
		<-s.tickChannel
	}()
	for {
		<-s.getTickChannel

	gxc:
		record, err := reader.Read()
		if err != nil {
			s.tickChannel <- s.tick
			return
		}

		layout := "2006-01-02 15:04:05.000Z"
		tx, err := time.Parse(layout, record[2])
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}

		bidVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		askVal, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		// Save data to database
		tick := models.Tick{
			Symbol:    record[1],
			Timestamp: tx.Unix(),
			Ask:       math.Round(askVal*100000) / 100000,
			Bid:       math.Round(bidVal*100000) / 100000,
		}

		// 过滤高点差数据
		if tick.Ask-tick.Bid > 0.0001 {
			goto gxc
		}

		if math.Abs(tick.Ask-s.tick.Ask) < 0.0003 {
			goto gxc
		}

		s.setTick(tick)
		s.tickChannel <- tick
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
	go func() {
		for {
			time.Sleep(time.Second)
			var account models.Account
			err := rs.Bb.Model(models.Account{}).First(&account).Error
			if err == nil {
				rs.Bb.Model(&models.AccountLog{}).Where("funding_dynamics = ?", account.FundingDynamics).FirstOrCreate(&models.AccountLog{
					Account:         account.Account,
					FundingDynamics: account.FundingDynamics,
				})
			}
		}
	}()
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

	tick2 := s.GetTick2()

	var myProfile float64
	var margin float64
	for idx, order := range orders {
		price := tick2.Ask
		var profile float64
		if order.Type == 1 { // 如果当前订单为sell
			price = tick2.Ask // 这做空平仓
			profile = math.Round(((order.Price-price)*order.Volume*100000)*100) / 100
		} else {
			price = tick2.Bid // 做多
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

var recordMu sync.Mutex
var recordLoss float64   // 最大浮亏
var recordProfit float64 // 最大浮盈
var lossNum int          // 亏损时持仓数量
var profitNum int        // 盈利时持仓数量
var lossAmount float64   // 亏损时手数
var profitAmount float64 // 盈利时手数
var comment string

func (s *Storage) Record(orders []models.RespOrderPosition) {
	var prf float64
	var vol float64
	for _, v := range orders {
		vol += v.Volume
		prf += v.Profit
	}

	recordMu.Lock()
	defer recordMu.Unlock()

	if prf < 0 {
		// 寻找最低
		if recordLoss > prf {
			recordLoss = prf
			lossNum = len(orders)
			lossAmount = vol
			if len(orders) > 0 {
				comment = orders[0].Comment
			}
		}
	}

	if prf > 0 {
		// 寻找新高
		if prf > recordProfit {
			recordProfit = prf
			profitNum = len(orders)
			profitAmount = vol
			if len(orders) > 0 {
				comment = orders[0].Comment
			}
		}
	}
}

func (s *Storage) RecordUp(timer int64) {
	// update
	s.Bb.Model(&models.OrderHistoryTick{}).
		Create(&models.OrderHistoryTick{
			Time:     timer,
			TimeStr:  time.Unix(timer, 0).Format("2006-01-02 15:04:05"),
			Profit:   recordLoss,
			Position: lossNum,
			Volume:   lossAmount,
			Comment:  comment,
		})

	s.Bb.Model(&models.OrderHistoryTick{}).
		Create(&models.OrderHistoryTick{
			Time:     timer,
			TimeStr:  time.Unix(timer, 0).Format("2006-01-02 15:04:05"),
			Profit:   recordProfit,
			Position: profitNum,
			Volume:   profitAmount,
			Comment:  comment,
		})

	recordMu.Lock()
	defer recordMu.Unlock()
	recordLoss = 0
	lossNum = 0
	lossAmount = 0
	recordProfit = 0
	profitNum = 0
	profitAmount = 0
	comment = ""
}
