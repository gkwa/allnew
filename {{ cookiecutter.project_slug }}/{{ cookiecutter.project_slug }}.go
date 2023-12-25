package {{ cookiecutter.project_slug }}

import (
	"fmt"
	"log/slog"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	LogLevel  string `short:"l" long:"log-level" choice:"debug" choice:"info" choice:"warn" choice:"error" default:"info" required:"false"`
	LogFormat string `long:"log-format" choice:"text" choice:"json" default:"text" required:"false"`
}

func Execute() int {
	_, err := flags.Parse(&opts)
	if err != nil {
		return 1
	}

	logger, err := getLogger(opts.LogLevel, opts.LogFormat)
	if err != nil {
		slog.Error("getLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	err = run()
	if err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}
	return 0
}

func run() error {
	slog.Debug("test", "LogLevel", opts.LogLevel)
	slog.Error("test", "test", "Error")

	fmt.Printf("LogLevel: %v\n", opts.LogLevel)

	return nil
}
