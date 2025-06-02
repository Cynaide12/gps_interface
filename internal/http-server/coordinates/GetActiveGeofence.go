package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	"gps_backend/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

func GetActiveGeofence(log *slog.Logger, coordinatesHandler CoordinatesHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "http-server.handlers.coordinates.GetActiveGeofence"
		log.With(slog.String("fn", fn))

		geofence, err := coordinatesHandler.GetActiveGeofence()
		if err != nil {
			log.Error("failed to get geofence", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get geofence"))
			return
		}

		render.JSON(w, r, GeofenceResponse{
			Response:  response.OK(),
			Geofence: geofence,
		})
	}
}
