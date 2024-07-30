package storage

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Storage struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewStorage(db *gorm.DB, cache *redis.Client) *Storage {

	//db.AutoMigrate(
	//	//&models.Account{},
	//	&models.Error{},
	//	&models.History{},
	//	//&models.Positions{},
	//	&models.Statistics{},
	//	&models.Strategy{},
	//	&models.TimeSeriesPosition{},
	//)

	return &Storage{db: db, cache: cache}
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}

func (s *Storage) Cache() *redis.Client {
	return s.cache
}
