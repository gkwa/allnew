package allnew

import (
	"log/slog"

	"github.com/taylormonacelli/allnew/logging"
	optmod "github.com/taylormonacelli/allnew/options"
)

func Main(options optmod.Options) int {
	logger := logging.GetLogger(options.LogLevel, options.LogFormat)

	slog.SetDefault(logger)

	slog.Debug("test", "test", "Debug")
	slog.Info("test", "test", "Info")
	return 0
}
