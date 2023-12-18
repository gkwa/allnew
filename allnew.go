package allnew

import (
	"log/slog"
)

func Main(logger *slog.Logger) int {
	slog.SetDefault(logger)
	slog.Debug("test", "test", "Debug")
	slog.Info("test", "test", "Info")
	return 0
}
