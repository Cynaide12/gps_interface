package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	"gps_backend/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

func GetCoordinates(log *slog.Logger, coordinatesHandler CoordinatesHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "http-server.handlers.project.getProject"
		log.With(slog.String("fn", fn))


		coordinates, err := coordinatesHandler.GetCoordinates()
		if err != nil {
			log.Error("failed to get coordinates", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to get coordinates"))
			return
		}

		render.JSON(w, r, CoordinatesResponse{
			Response:    response.OK(),
			Coordinates: coordinates,
		})
	}
}
