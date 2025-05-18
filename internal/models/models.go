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
	ID        uint      `gorm:"primaryKey" json:"id"`
	DeviceID  string    `gorm:"size:36;not null" json:"device_id"` // Идентификатор устройства
	Latitude  float64   `gorm:"type:decimal(10,6);not null" json:"latitude"`
	Longitude float64   `gorm:"type:decimal(10,6);not null" json:"longitude"`
	Speed     float64   `gorm:"type:decimal(6,2)" json:"speed"`    // км/ч
	Altitude  float64   `gorm:"type:decimal(6,2)" json:"altitude"` // метры
	Timestamp time.Time `gorm:"index;not null" json:"timestamp"`    // Временная метка
	Satellites int      `json:"satellites"`                        // Кол-во спутников
}

type Device struct {
	ID          string    `gorm:"primaryKey;size:36" json:"id"`
	Name        string    `gorm:"size:100" json:"name"`
	LastSeen    time.Time `json:"last_seen"`
	BatteryLevel int      `json:"battery_level"` // Опционально
}