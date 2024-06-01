package logger

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

var Log = zap.NewNop() //nolint:gochecknoglobals // example

func Initialize(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return fmt.Errorf("failed to parse log level: %w", err)
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	zapLogger, err := cfg.Build()
	if err != nil {
		return fmt.Errorf("failed to build zap logger: %w", err)
	}

	Log = zapLogger

	return nil
}

func RequestLogger(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log.Debug("got incoming HTTP request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
		)

		next(w, r)
	})
}
