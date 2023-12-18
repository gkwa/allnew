package {{ cookiecutter.project_slug }}

import (
	"flag"
	"log/slog"
)

func Execute() int {
	options := Options{}
	flag.StringVar(&options.LogLevel, "log-level", "info", "Log level (debug, info, warn, error), defult: info")
	flag.StringVar(&options.LogFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	logger, err := getLogger(options.LogLevel, options.LogFormat)
	if err != nil {
		slog.Error("getLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	slog.Debug("test", "test", "Debug")
	slog.Info("test", "test", "Info")
	slog.Error("test", "test", "Error")
	return 0
}
