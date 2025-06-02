package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	"gps_backend/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

func GetGeofences(log *slog.Logger, coordinatesHandler CoordinatesHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "http-server.handlers.coordinates.GetGeofences"
		log.With(slog.String("fn", fn))

		geofences, err := coordinatesHandler.GetGeofences()
		if err != nil {
			log.Error("failed to get geofences", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get geofences"))
			return
		}

		render.JSON(w, r, GeofencesResponse{
			Response:  response.OK(),
			Geofences: geofences,
		})
	}
}
