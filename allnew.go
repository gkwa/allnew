package allnew

import (
	"log/slog"

	"github.com/taylormonacelli/allnew/logging"
	optmod "github.com/taylormonacelli/allnew/options"
)

func Main(options optmod.Options) int {
	logger, err := logging.GetLogger(options.LogLevel, options.LogFormat)
	if err != nil {
		slog.Error("GetLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	slog.Debug("test", "test", "Debug")
	slog.Info("test", "test", "Info")
	return 0
}
