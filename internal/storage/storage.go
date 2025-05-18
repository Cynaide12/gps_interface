package storage

import (
	"database/sql"
	"fmt"
	"gps_backend/internal/config"
	"gps_backend/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

func New(cfg *config.Config) (*Storage, error) {
	const fn = "http-server.storage.New"

	dbPath := cfg.DBServer.DBPath
	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	// миграции
	gormDB.AutoMigrate(
		&models.Coordinate{},
		&models.Device{},
	)

	return &Storage{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}, nil
}

func (s *Storage) GetCoordinates() ([]models.Coordinate, error) {
	const fn = "internal/storage/GetCoordinates"

	var coordinates []models.Coordinate
	if err := s.gormDB.Model(&models.Coordinate{}).Find(&coordinates).Error; err != nil {
		return nil, fmt.Errorf("%s: %s", fn, err)
	}
	return coordinates, nil
}
