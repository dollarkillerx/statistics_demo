package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {

	//db.AutoMigrate()

	return &Storage{db: db}
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}
