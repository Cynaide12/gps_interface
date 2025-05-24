package handlers_coordinates

import (
	response "gps_backend/internal/lib/api"
	"gps_backend/internal/lib/logger/sl"
	"gps_backend/internal/models"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func AddCoordinate(log *slog.Logger, coordinatesHandler CoordinatesHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "http-server.handlers.project.getProject"
		log.With(slog.String("fn", fn))

		var req models.Coordinate

		// декодируем запрос в структуру
		if err := render.Decode(r, &req); err != nil {
			log.Error("failed to decode request", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.Error("invalid request"))
			return
		}

		// валидация запроса
		if err := response.ValidateRequest(&req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.ValidationError(validateErr))
			return
		}

		newCoordinate, err := coordinatesHandler.AddCoordinate(req)
		if err != nil{
			log.Error("failed to add coordinate", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.Error("failed to add coordinate"))
			return
		}
		

		render.JSON(w, r, CoordinateResponse{
			Response:     response.OK(),
			Coordinates: *newCoordinate,
		})
	}
}
