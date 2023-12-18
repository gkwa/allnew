package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/taylormonacelli/allnew"
	"github.com/taylormonacelli/littlecow"
)

var (
	logLevelString string
	logFormat      string
)

func main() {
	flag.StringVar(&logLevelString, "log-level", "info", "Log level (debug, info, warn, error), defult: info")
	flag.StringVar(&logFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	logLevel, err := littlecow.LevelFromString(logLevelString)
	if err != nil {
		slog.Error("LevelFromString", "error", err)
		return
	}

	opts := littlecow.NewHandlerOptions(logLevel, littlecow.RemoveTimestampAndTruncateSource)

	var handler slog.Handler
	handler = slog.NewTextHandler(os.Stderr, opts)
	if logFormat == "json" {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	}

	logger := slog.New(handler)

	code := allnew.Main(logger)
	os.Exit(code)
}
