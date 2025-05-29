package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	"gps_backend/internal/models"
)

type CoordinatesHandler interface {
	GetCoordinates() ([]models.Coordinate, error)
	GetLastCoordinates() (*models.Coordinate, error) 
	AddCoordinate(coordinate models.Coordinate) (*models.Coordinate, error)
}

type CoordinateResponse struct {
	response.Response
	Coordinates *models.Coordinate
}

type CoordinatesResponse struct {
	response.Response
	Coordinates []models.Coordinate
}
