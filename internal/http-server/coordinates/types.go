package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	"gps_backend/internal/models"
)

type CoordinatesHandler interface {
	GetCoordinates() ([]models.Coordinate, error)
}


type CoordinatesResponse struct{
	response.Response
	Coordinates models.Coordinate
}