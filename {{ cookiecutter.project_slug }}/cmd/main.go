package main

import (
	"flag"
	"os"

	"github.com/taylormonacelli/{{ cookiecutter.project_slug }}"
	optmod "github.com/taylormonacelli/{{ cookiecutter.project_slug }}/options"
)

func main() {
	options := optmod.Options{}
	flag.StringVar(&options.LogLevel, "log-level", "info", "Log level (debug, info, warn, error), defult: info")
	flag.StringVar(&options.LogFormat, "log-format", "", "Log format (text or json)")

	flag.Parse()

	code := {{ cookiecutter.project_slug }}.Main(options)
	os.Exit(code)
}
