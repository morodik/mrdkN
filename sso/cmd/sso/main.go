package main

import (
	"log/slog"
	"os"

	"example.com/grpc-service/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//иницилизировать объект конфига
	cfg := config.MustLoad()

	//инициализировать объект логирования, логгер
	log := setupLogger(cfg.Env)

	log.Info("starting aplication", slog.String("env", cfg.Env))
	//инициализация приложения (app)
	//запуск gRPC

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch envLocal {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
