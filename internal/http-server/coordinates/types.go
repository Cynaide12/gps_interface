package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	"gps_backend/internal/models"
)

type CoordinatesHandler interface {
	GetCoordinates() ([]models.Coordinate, error)
	GetLastCoordinates() (*models.Coordinate, error)
	AddCoordinate(coordinate models.Coordinate) (*models.Coordinate, error)
	GetActiveGeofence() (*models.Geofence, error)
	GetGeofences() ([]models.Geofence, error)
	AddGeofence(geofence models.Geofence) (*models.Geofence, error)
	UpdateGeofence(geofence models.Geofence) (*models.Geofence, error)
	DeleteGeofence(geofence models.Geofence) error
	SetActiveGeofence(geofence models.Geofence) error
	GetConnectionQuality() string 
}

type CoordinateResponse struct {
	response.Response
	Coordinates      *models.Coordinate
	IsInsideGeofence bool `json:"is_inside_geofence"`
	DeviceStatus         string `json:"device_status"`
}

type CoordinatesResponse struct {
	response.Response
	Coordinates []models.Coordinate
}

type GeofencesResponse struct {
	response.Response
	Geofences []models.Geofence
}

type GeofenceResponse struct {
	response.Response
	Geofence *models.Geofence
}
