package logging

import (
	"log/slog"
	"os"

	"github.com/taylormonacelli/littlecow"
)

func GetLogger(logLevelString, logFormat string) *slog.Logger {
	logLevel, err := littlecow.LevelFromString(logLevelString)
	if err != nil {
		slog.Error("LevelFromString", "error", err)
		return nil
	}

	opts := littlecow.NewHandlerOptions(logLevel, littlecow.RemoveTimestampAndTruncateSource)

	var handler slog.Handler
	handler = slog.NewTextHandler(os.Stderr, opts)
	if logFormat == "json" {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}

	return slog.New(handler)
}
