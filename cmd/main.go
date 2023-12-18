package main

import (
	"flag"
	"os"

	"github.com/taylormonacelli/allnew"
	optmod "github.com/taylormonacelli/allnew/options"
)

func main() {
	options := optmod.Options{}
	flag.StringVar(&options.LogLevel, "log-level", "info", "Log level (debug, info, warn, error), defult: info")
	flag.StringVar(&options.LogFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	code := allnew.Main(options)
	os.Exit(code)
}
