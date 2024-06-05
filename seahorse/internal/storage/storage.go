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
	for {
		<-s.getTickChannel
		record, err := reader.Read()
		if err != nil {
			s.tickChannel <- models.Tick{Over: true}
			s.setTick(models.Tick{Over: true})
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
			Ask:       math.Round(bidVal*100000) / 100000,
			Bid:       math.Round(askVal*100000) / 100000,
		}
		s.tickChannel <- tick

		s.setTick(tick)
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

	db.AutoMigrate(&models.Order{}, &models.Account{})

	rs := &Storage{Bb: db}

	go rs.heartbeat()
	return rs
}
