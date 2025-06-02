package models

import (
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrInvalidModel   = errors.New("invalid model")
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Coordinate struct {
	ID         uint      `gorm:"primaryKey"`
	DeviceID   string    `gorm:"size:36;not null"` // Идентификатор устройства
	Latitude   float64   `gorm:"type:decimal(10,6);not null"`
	Longitude  float64   `gorm:"type:decimal(10,6);not null"`
	Speed      float64   `gorm:"type:decimal(6,2)"` // км/ч
	Altitude   float64   `gorm:"type:decimal(6,2)"` // метры
	Timestamp  time.Time `gorm:"index;not null"`    // Временная метка
	Satellites int       `json:"satellites"`        // Кол-во спутников
}

type Device struct {
	ID           string `gorm:"primaryKey;size:36"`
	Name         string `gorm:"size:100"`
	LastSeen     time.Time
	BatteryLevel int // Опционально
}

type Geofence struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:100;not null"`
	Latitude  float64 `gorm:"type:decimal(10,6);not null"`
	Longitude float64 `gorm:"type:decimal(10,6);not null"`
	Radius    float64 `gorm:"not null"` // Радиус в метрах
	IsActive  bool    `gorm:"not null"`
	CreatedAt time.Time
}
