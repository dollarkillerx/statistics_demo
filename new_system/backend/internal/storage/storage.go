package storage

import (
	"github.com/dollarkillerx/backend/pkg/models"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {

	db.AutoMigrate(
		&models.Account{},
		&models.Error{},
		&models.History{},
		&models.Positions{},
		&models.Statistics{},
		&models.Strategy{},
		&models.TimeSeriesPosition{},
	)

	return &Storage{db: db}
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}
