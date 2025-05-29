package main

import (
	"gps_backend/internal/config"
	handlers_coordinates "gps_backend/internal/http-server/coordinates"
	"gps_backend/internal/http-server/logger"
	slogpretty "gps_backend/internal/lib/logger/handlers"
	"gps_backend/internal/lib/logger/sl"
	"gps_backend/internal/storage"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting gps-backend", slog.String("env", cfg.Env))

	log.Debug("debug messages are enabled")

	storage := initDB(cfg, log)

	initRouter(cfg, log, &storage)
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log

}

func initRouter(cfg *config.Config, log *slog.Logger, storage *storage.Storage) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(logger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

		r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*", "https://47fc-185-77-216-6.ngrok-free.app"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		// MaxAge: 300,
		Debug: true,
	}))


	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	r.Get("/api/get_coordinates", handlers_coordinates.GetCoordinates(log, storage))
	r.Get("/api/get_last_coordinates", handlers_coordinates.GetLastCoordinates(log, storage))
	r.Post("/add_coordinate", handlers_coordinates.AddCoordinate(log, storage))

	log.Info("starting server", slog.String("address", srv.Addr))

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", sl.Err(err))

		os.Exit(1)
	}

	log.Error("server stopped")

}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func initDB(cfg *config.Config, log *slog.Logger) storage.Storage {
	storage, err := storage.New(cfg)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	return *storage
}
