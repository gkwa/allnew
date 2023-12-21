package {{ cookiecutter.project_slug }}

import (
	"fmt"
	"flag"
	"log/slog"
)

type Options struct {
	LogFormat string
	LogLevel  string
}


func Execute() int {
	options := parseArgs()

	logger, err := getLogger(options.LogLevel, options.LogFormat)
	if err != nil {
		slog.Error("getLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	err = run(options)
	if err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}
	return 0
}

func parseArgs() Options {
	const logLevelHelp = "Log level (debug, info, warn, error), default: info"

	options := Options{}

	ll := flag.String("ll", "info", fmt.Sprintf("%s (shorthand)", logLevelHelp))
	flag.StringVar(&options.LogLevel, "log-level", *ll, logLevelHelp)

	flag.StringVar(&options.LogFormat, "log-format", "text", "Log format (text or json)")

	flag.Parse()

	return options
}

func run(options Options) error {
	slog.Debug("test", "test", "Debug")
	slog.Debug("test", "LogLevel", options.LogLevel)
	slog.Info("test", "test", "Info")
	slog.Error("test", "test", "Error")

	return nil
}