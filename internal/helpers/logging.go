package helpers

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func SetupLogging() {
	if os.Getenv("LOG_JSON") == "true" {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	} else {
		logger := slog.New(tint.NewHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	}
}
