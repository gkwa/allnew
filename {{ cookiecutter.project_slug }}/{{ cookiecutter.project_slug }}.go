package {{ cookiecutter.project_slug }}

import (
	"log/slog"

	optmod "github.com/taylormonacelli/{{ cookiecutter.project_slug }}/options"
)

func Main(options optmod.Options) int {
	logger, err := getLogger(options.LogLevel, options.LogFormat)
	if err != nil {
		slog.Error("GetLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	slog.Debug("test", "test", "Debug")
	slog.Info("test", "test", "Info")
	slog.Error("test", "test", "Error")
	return 0
}
