package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	lib_geofence "gps_backend/internal/lib/geofence"
	"gps_backend/internal/lib/logger/sl"
	"gps_backend/internal/models"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

func GetLastCoordinates(log *slog.Logger, coordinatesHandler CoordinatesHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "http-server.handlers.coordinates.GetLastCoordinates"
		log.With(slog.String("fn", fn))

		coordinates, err := coordinatesHandler.GetLastCoordinates()
		if err != nil && err != models.ErrRecordNotFound {
			log.Error("failed to get coordinates", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get coordinates"))
			return
		}

		activeGeofence, err := coordinatesHandler.GetActiveGeofence()
		if err != nil {
			if err == models.ErrRecordNotFound {
				log.Error("failed to get active geofence", sl.Err(err))
				w.WriteHeader(http.StatusInternalServerError)
				render.JSON(w, r, response.Error("failed to get active geofence"))
				return
			}
			log.Error("failed to get active geofence", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get active geofence"))
			return
		}

		deviceStatus := coordinatesHandler.IsDeviceOnline()


		IsInsideGeofence := lib_geofence.IsInsideGeofence(coordinates.Latitude,
			coordinates.Longitude,
			activeGeofence.Latitude,
			activeGeofence.Longitude,
			activeGeofence.Radius)

		render.JSON(w, r, CoordinateResponse{
			Response:         response.OK(),
			Coordinates:      coordinates,
			IsInsideGeofence: IsInsideGeofence,
			IsOnline: deviceStatus,
		})
	}
}
