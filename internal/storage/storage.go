package storage

import (
	"database/sql"
	"fmt"
	"gps_backend/internal/config"
	"gps_backend/internal/models"
	"time"

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
		&models.Geofence{},
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

func (s *Storage) GetLastCoordinates() (*models.Coordinate, error) {
	const fn = "internal/storage/GetLastCoordinates"

	var coordinates *models.Coordinate
	if err := s.gormDB.Model(&models.Coordinate{}).Order("id DESC").First(&coordinates).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrRecordNotFound
		}
		return nil, fmt.Errorf("%s: %s", fn, err)
	}
	return coordinates, nil
}

func (s *Storage) AddCoordinate(coordinate models.Coordinate) (*models.Coordinate, error) {
	const fn = "internal/storage/AddCoordinate"

	if err := s.gormDB.Create(&coordinate).Error; err != nil {
		return nil, fmt.Errorf("%s: %s", fn, err)
	}
	return &coordinate, nil
}

func (s *Storage) AddGeofence(geofence models.Geofence) (*models.Geofence, error) {
	const fn = "internal/storage/AddGeofence"

	if err := s.gormDB.Create(&geofence).Error; err != nil {
		return nil, fmt.Errorf("%s: %s", fn, err)
	}
	return &geofence, nil
}

func (s *Storage) UpdateGeofence(geofence models.Geofence) (*models.Geofence, error) {
	const fn = "internal/storage/UpdateGeofence"

	if err := s.gormDB.Model(&models.Geofence{}).Where("id = ?", geofence.ID).Updates(&geofence).Error; err != nil {
		return nil, fmt.Errorf("%s: %s", fn, err)
	}
	return &geofence, nil
}

func (s *Storage) SetActiveGeofence(geofence models.Geofence) error {
	const fn = "internal/storage/SetActiveGeofence"

	if err := s.gormDB.Model(&models.Geofence{}).Where("is_active = ?", true).Update("is_active", false).Error; err != nil {
		return fmt.Errorf("%s: %s", fn, err)
	}

	if err := s.gormDB.Model(&models.Geofence{}).Where("id = ?", geofence.ID).Update("is_active", true).Error; err != nil {
		return fmt.Errorf("%s: %s", fn, err)
	}
	return nil
}

func (s *Storage) DeleteGeofence(geofence models.Geofence) error {
	const fn = "internal/storage/DeleteGeofence"

	if err := s.gormDB.Delete(&geofence).Error; err != nil {
		return fmt.Errorf("%s: %s", fn, err)
	}
	return nil
}

func (s *Storage) GetGeofences() ([]models.Geofence, error) {
	const fn = "internal/storage/GetGeofences"

	var geofences []models.Geofence
	if err := s.gormDB.Model(&models.Geofence{}).Find(&geofences).Error; err != nil {
		return nil, fmt.Errorf("%s: %s", fn, err)
	}
	return geofences, nil
}

func (s Storage) GetActiveGeofence() (*models.Geofence, error) {
	const fn = "internal/storage/GetActiveGeofence"

	var geofence *models.Geofence
	if err := s.gormDB.Model(&models.Geofence{}).Where("is_active = ?", true).First(&geofence).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrRecordNotFound
		}
		return nil, fmt.Errorf("%s: %s", fn, err)
	}
	return geofence, nil
}

func (s Storage) IsDeviceOnline() bool {
	var lastPos models.Coordinate
	err := s.gormDB.
		Order("created_at DESC").
		First(&lastPos).Error

	if err != nil {
		return false
	}

	// Считаем устройство онлайн, если координата обновлялась
	// в последние 30 секунд (15*2 с запасом)
	return time.Since(lastPos.CreatedAt) <= 50*time.Second
}
